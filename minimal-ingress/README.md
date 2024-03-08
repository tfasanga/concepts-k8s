# Minikube

Run minikube with socket_vmnet network:

```shell
brew install qemu
brew install socket_vmnet
brew tap homebrew/services
HOMEBREW=$(which brew) && sudo ${HOMEBREW} services start socket_vmnet
```

If minikube was running before then stop & delete it:

```shell
minikube stop
minikube delete
```

```shell
minikube start --network socket_vmnet
```

or

```shell
minikube start --driver qemu --network socket_vmnet
```

```shell
kubectx minikube
```

# Kind

```shell
brew install kind
kind create cluster
```

```shell
kubectx kind-kind
```

# Download Nginx

```shell
mkdir nginx
wget https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.10.0/deploy/static/provider/baremetal/deploy.yaml -O nginx/deploy.yaml
```

# Prerequisite: Install Nginx Ingress Controller in Kubernetes

```shell
kubectl apply -f nginx/deploy.yaml
```

# Build

Build and Save Docker image:

```shell
podman build -t my-go-api-ingress:1.0 .
podman save my-go-api-ingress:1.0 -o my-go-api-ingress-image.tar
```

# Load docker image to Minikube

```shell
minikube image load my-go-api-ingress-image.tar
```

# Load docker image to Kind

```shell
kind load image-archive my-go-api-ingress-image.tar
```

# Test run in Podman

```shell
podman run --rm -p 8080:8080 --name my-go-api-ingress my-go-api-ingress:1.0 
```

## Remove image from minikube

```shell
minikube image rm docker.io/localhost/my-go-api-ingress:1.0
```

## Remove image from podman

```shell
podman image rm my-go-api-ingress:1.0
```

# Base

[TEST_BASE.md](TEST_BASE.md)

# Dev Overlay

[TEST_OVERLAY_DEV.md](TEST_OVERLAY_DEV.md)

# Remove Nginx Ingress Controller from Kubernetes

```shell
kubectl delete -f nginx/deploy.yaml
```



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

