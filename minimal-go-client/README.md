# Go

```shell
go get github.com/anthdm/hollywood
go get k8s.io/apimachinery
go get k8s.io/api
go get k8s.io/client-go
go get k8s.io/klog/v2
go get github.com/spf13/cobra
go get github.com/spf13/viper
```

If you have RBAC enabled on your cluster,
run the following snippet to create role binding
which will grant the default service account view permissions.

```shell
kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=default:default
```

If your application runs in a Pod in the cluster:
https://github.com/kubernetes/client-go/blob/v0.29.2/examples/in-cluster-client-configuration


# Minikube

Run minikube with socket_vmnet network:

```shell
brew install qemu
brew install socket_vmnet
brew tap homebrew/services
HOMEBREW=$(which brew) && sudo ${HOMEBREW} services start socket_vmnet
```

If minikube was running before, then stop & delete it:

```shell
minikube stop
minikube delete
```

```shell
minikube start --network socket_vmnet
```

```shell
kubectx minikube
```

# Enable local insecure registry in Minikube 

```shell
minikube addons enable registry
```

```shell
curl http://$(minikube ip):5000/v2/
```

returns `{}` 

Determine Minikube IP:

```shell
minikube ip
```

returns `192.168.105.5`

*Note: On Mac OS Podman runs inside a VM.*

Enter Podman VM:

```shell
podman machine ssh
```

```
sudo vi /etc/containers/registries.conf
```

Add following section:

```
[[registry]]
location = "192.168.105.5:5000"
insecure = true
```

```
# Exit Podman VM
exit
```

Restart Podman VM:

```shell
podman machine stop
podman machine start
```

# Kind

```shell
brew install kind
kind create cluster
```

```shell
kubectx kind-kind
```

# Build

Build and Save Docker image:

```shell
podman build -t my-go-api-goclient:1.0 .
podman save my-go-api-goclient:1.0 -o my-go-api-goclient-image.tar
```
```shell
minikube image ls
```

Build and Push Docker image to local registry:

```shell
podman build -t $(minikube ip):5000/my-go-api-goclient:1.0 .
podman push $(minikube ip):5000/my-go-api-goclient:1.0
```

```shell
curl -X GET http://$(minikube ip):5000/v2/_catalog
```

```
{"repositories":["my-go-api-goclient"]}
```

```shell
curl -X GET http://$(minikube ip):5000/v2/my-go-api-goclient/tags/list
```

```
{"name":"my-go-api-goclient","tags":["1.0"]}
```


# Load docker image to Minikube

```shell
minikube image load my-go-api-goclient-image.tar
```

# Load docker image to Kind

```shell
kind load image-archive my-go-api-goclient-image.tar
```

## Remove image from minikube

```shell
minikube image rm docker.io/localhost/my-go-api-goclient:1.0
```

## Remove image from podman

```shell
podman image rm my-go-api-goclient:1.0
```

# In Cluster

[TEST_IN_CLUSTER.md](TEST_IN_CLUSTER.md)

# Out of Cluster (on host)

[TEST_OUT_OF_CLUSTER.md](TEST_OUT_OF_CLUSTER.md)

# Notes

Show logs:

```shell
kubectl logs -l app=go-api-label -f
```

Show pods:

```shell
kubectl get pods
```

Show service:

```shell
kubectl describe service go-api-service
kubectl get service
```

