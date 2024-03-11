package main

import (
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/anthdm/hollywood/actor"
	"github.com/spf13/viper"
	"io"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"k8s.io/client-go/tools/clientcmd"
	"k8s.io/client-go/util/homedir"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"runtime"
	"time"
)

type programContext struct {
	clientset *kubernetes.Clientset
	config    *Config
}

func (pc *programContext) httpHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Reading file")
	content, err := os.ReadFile("content/content.txt")
	if err != nil {
		fmt.Fprintf(w, "Cannot read content file\n")
	} else {
		fmt.Fprintf(w, string(content))
	}

	err = printNumberOfPods(w, pc.clientset, pc.config.namespaces)
	if err != nil {
		log.Println(err.Error())
	}
}

func startListener(pc *programContext) error {
	log.Println("Starting listener 2")
	http.HandleFunc("/", pc.httpHandler)
	return http.ListenAndServe(":8080", nil)
}

type Config struct {
	namespaces []string
}

func NewConfig(v *viper.Viper) (cfg *Config, err error) {
	cfg = &Config{
		namespaces: v.GetStringSlice("namespaces"),
	}

	return cfg, nil
}

func main() {
	cfgFile := flag.String("config", "", "config file")
	inCluster := flag.Bool("pod", false, "run in POD")
	kubeconfig := flag.String("kubeconfig", "", "absolute path to the kubeconfig file")
	flag.Parse()

	if *cfgFile != "" {
		// Use config file from the flag.
		fmt.Fprintf(os.Stdout, "Config file: %s", *cfgFile)
		viper.SetConfigFile(*cfgFile)
	} else {
		fmt.Fprintf(os.Stdout, "No Config file specified\n")
	}

	// If a config file is found, read it in.
	if *cfgFile != "" {
		if err := viper.ReadInConfig(); err != nil {
			fmt.Fprintln(os.Stderr, "Failed to read config file: ", viper.ConfigFileUsed())
		} else {
			fmt.Fprintln(os.Stdout, "Config file read done: ", viper.ConfigFileUsed())
		}
	}

	cfg, err := NewConfig(viper.GetViper())
	if err != nil {
		panic(err.Error())
	}

	var restConfig *rest.Config

	if *inCluster {
		fmt.Println("Running in Kubernetes Pod")
		// https://github.com/kubernetes/client-go/blob/v0.29.2/examples/in-cluster-client-configuration/main.go
		// creates the in-cluster config
		restConfig, err = rest.InClusterConfig()
	} else {
		if *kubeconfig == "" {
			if home := homedir.HomeDir(); home != "" {
				*kubeconfig = filepath.Join(home, ".kube", "config")
			}
		}

		restConfig, err = clientcmd.BuildConfigFromFlags("", *kubeconfig)
	}

	if err != nil {
		panic(err.Error())
	}

	clientSet, err := kubernetes.NewForConfig(restConfig)
	if err != nil {
		panic(err.Error())
	}

	pc := &programContext{clientset: clientSet, config: cfg}

	err = printNumberOfPods(os.Stdout, pc.clientset, pc.config.namespaces)
	if err != nil {
		panic(err.Error())
	}

	// create actor engine
	engine, err := actor.NewEngine(actor.NewEngineConfig())
	var pid *actor.PID = engine.Spawn(newKubeWatcherActor, "kube-watcher")
	engine.Send(pid, "hello world!")

	// NewSharedInformerFactory will create a new ShareInformerFactory for "all namespaces”
	// 30*time.Second is the re-sync period to update the in-memory cache of informer
	informerFactory := informers.NewSharedInformerFactory(clientSet, 30*time.Second)

	podInformer := informerFactory.Core().V1().Pods()

	c := Controller{engine, pid}

	informer := podInformer.Informer()

	_, err = informer.AddEventHandler(
		&cache.ResourceEventHandlerFuncs{
			AddFunc:    c.HandleAdd,
			DeleteFunc: c.HandleDelete,
			UpdateFunc: c.HandleUpdate,
		})
	if err != nil {
		panic(err.Error())
	}

	// creating a unbuffered channel to synchronize the update
	stopChan := make(chan struct{})

	// To stop the channel automatically at the end of our main functions
	defer close(stopChan)

	go podInformer.Informer().Run(stopChan)

	err = startListener(pc)
	if err != nil {
		panic(err.Error())
	}
}

func printNumberOfPods(w io.Writer, clientset *kubernetes.Clientset, namespaces []string) error {
	if len(namespaces) == 0 {
		// get pods in all the namespaces by omitting namespace
		pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
		if err != nil {
			return err
		}
		fmt.Fprintf(w, "There are %d pods in all namespaces in the cluster\n", len(pods.Items))

	} else {
		for _, namespace := range namespaces {
			pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				return err
			}
			fmt.Fprintf(w, "There are %d pods in namespace %s in the cluster\n", len(pods.Items), namespace)
		}
	}

	return nil
}

type message struct{ data string }

type MsgAdded struct {
	AddedObj interface{}
}

type MsgDeleted struct {
	DeletedObj interface{}
}

type MsgUpdated struct {
	OldObj interface{}
	NewObj interface{}
}

type podWatcher struct{}

func newKubeWatcherActor() actor.Receiver {
	return &podWatcher{}
}

func (h *podWatcher) Receive(ctx *actor.Context) {
	switch msg := ctx.Message().(type) {
	case actor.Initialized:
		fmt.Println("kube-watcher has initialized")
	case actor.Started:
		fmt.Println("kube-watcher has started")
	case actor.Stopped:
		fmt.Println("kube-watcher has stopped")
	case MsgAdded:
		//fmt.Printf("Added: %s", msg.AddedObj)
		fmt.Printf("Added %T %s\n", msg.AddedObj, GoId())
		if pod, ok := (msg.AddedObj).(*v1.Pod); ok {
			fmt.Printf("POD Name: %s/%s\n", pod.Namespace, pod.Name)
		}
	case MsgDeleted:
		//fmt.Printf("Deleted: %s", msg.DeletedObj)
		fmt.Printf("Deleted %T %s\n", msg.DeletedObj, GoId())
		if pod, ok := (msg.DeletedObj).(*v1.Pod); ok {
			fmt.Printf("POD Name: %s/%s\n", pod.Namespace, pod.Name)
		}
	case MsgUpdated:
		//fmt.Printf("Updated: %s -> %s", msg.OldObj, msg.NewObj)
		//fmt.Printf("Updated %T -> %T %s\n", msg.OldObj, msg.NewObj, GoId())
		if msg.OldObj != msg.NewObj {
			fmt.Printf("POD Changed\n")
			if pod, ok := (msg.OldObj).(*v1.Pod); ok {
				fmt.Printf("Old POD Name: %s/%s\n", pod.Namespace, pod.Name)
			}
			if pod, ok := (msg.NewObj).(*v1.Pod); ok {
				fmt.Printf("New POD Name: %s/%s\n", pod.Namespace, pod.Name)
			}
		}
	case *message:
		fmt.Printf("received message: %s\n", msg.data)
	}
}

func GoId() string {
	var buffer [31]byte
	_ = runtime.Stack(buffer[:], false)
	return string(bytes.Fields(buffer[10:])[0])
}

func (c *Controller) HandleAdd(obj interface{}) {
	//fmt.Printf("inside add function %s\n", GoId())
	c.engine.Send(c.pid, MsgAdded{AddedObj: obj})
}

func (c *Controller) HandleDelete(obj interface{}) {
	//fmt.Printf("inside delete function %s\n", GoId())
	c.engine.Send(c.pid, MsgDeleted{DeletedObj: obj})
}

func (c *Controller) HandleUpdate(old, new interface{}) {
	//fmt.Printf("inside update function %s\n", GoId())
	c.engine.Send(c.pid, MsgUpdated{OldObj: old, NewObj: new})
}

type Controller struct {
	engine *actor.Engine
	pid    *actor.PID
}
