package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/anthdm/hollywood/actor"
	"io"
	v1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
	"log"
	"net/http"
	"os"
	"runtime"
	"time"
)

type programContext struct {
	clientset *kubernetes.Clientset
}

func (pc *programContext) httpHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Reading file")
	content, err := os.ReadFile("content/content.txt")
	if err != nil {
		fmt.Fprintf(w, "Cannot read content file\n")
	} else {
		fmt.Fprintf(w, string(content))
	}

	err = printNumberOfPods(w, pc.clientset)
	if err != nil {
		log.Println(err.Error())
	}
}

func startListener(pc *programContext) error {
	log.Println("Starting listener")
	http.HandleFunc("/", pc.httpHandler)
	return http.ListenAndServe(":8080", nil)
}

func main() {
	// https://github.com/kubernetes/client-go/blob/v0.29.2/examples/in-cluster-client-configuration/main.go
	// creates the in-cluster config
	config, err := rest.InClusterConfig()
	if err != nil {
		panic(err.Error())
	}

	clientSet, err := kubernetes.NewForConfig(config)
	if err != nil {
		panic(err.Error())
	}

	pc := &programContext{clientset: clientSet}

	err = printNumberOfPods(os.Stdout, pc.clientset)
	if err != nil {
		panic(err.Error())
	}

	engine, err := actor.NewEngine(actor.NewEngineConfig())
	var pid *actor.PID = engine.Spawn(newKubeWatcherActor, "kube-watcher")
	engine.Send(pid, "hello world!")

	// NewSharedInformerFactory will create a new ShareInformerFactory for "all namespaces”
	// 30*time.Second is the re-sync period to update the in-memory cache of informer
	informerFactory := informers.NewSharedInformerFactory(clientSet, 30*time.Second)

	podInformer := informerFactory.Core().V1().Pods()

	c := Controller{engine, pid}

	_, err = podInformer.Informer().AddEventHandler(
		&cache.ResourceEventHandlerFuncs{
			AddFunc:    c.HandleAdd,
			DeleteFunc: c.HandleDelete,
			UpdateFunc: c.HandleUpdate,
		})
	if err != nil {
		panic(err.Error())
	}

	// creating a unbuffered channel to synchronized the update
	stopChan := make(chan struct{})

	// To stop the channel automatically at the end of our main functions
	defer close(stopChan)

	go podInformer.Informer().Run(stopChan)

	err = startListener(pc)
	if err != nil {
		panic(err.Error())
	}
}

func printNumberOfPods(w io.Writer, clientset *kubernetes.Clientset) error {
	// get pods in all the namespaces by omitting namespace
	// Or specify namespace to get pods in particular namespace
	//pods, err := clientset.CoreV1().Pods("").List(context.TODO(), metav1.ListOptions{})
	pods, err := clientset.CoreV1().Pods("default").List(context.TODO(), metav1.ListOptions{})
	if err != nil {
		return err
	}
	_, err = fmt.Fprintf(w, "There are %d pods in the cluster\n", len(pods.Items))
	if err != nil {
		return err
	}
	return nil
}

type podWatcher struct{}

func newKubeWatcherActor() actor.Receiver {
	return &podWatcher{}
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
