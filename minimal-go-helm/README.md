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

# Build

Build Docker image:

```shell
podman build -t my-go-api-helm:1.0 .
```
```shell
podman save my-go-api-helm:1.0 -o my-go-api-helm-image.tar
```

Test run in Podman:

```shell
podman run --rm -p 8080:8080 --name my-go-api-helm my-go-api-helm:1.0 
```

# Load docker image to Minikube

```shell
minikube image load my-go-api-helm-image.tar
```

# Load docker image to Kind

```shell
kind load image-archive my-go-api-helm-image.tar
```

# Install in Kubernetes

```shell
helm install go-api helm --values helm/values.yaml
```

Logs:

```shell
kubectl logs -l app=go-api-label -f
```

Port forward:

```shell
export POD_NAME=$(kubectl get pods -l "app=go-api-label" -o jsonpath="{.items[0].metadata.name}")
echo "$POD_NAME"
kubectl port-forward $POD_NAME 8080:8080
```

Test:

```shell
curl http://127.0.0.1:8080
```

Install Dev Environment using NodePort:

```shell
helm install go-api helm --values env/dev-values.yaml
```

```shell
minikube ssh
```

inside minikube ssh;

```shell
curl http://10.0.2.15:30090
curl http://127.0.0.1:30090
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

Get the service URL from host:

```shell
minikube service go-api-service --url
```

# Uninstall Helm chart

```shell
helm uninstall go-api
```

## Remove image from minikube

```shell
minikube image rm docker.io/localhost/my-go-api-helm:1.0
```

## Remove image from podman

```shell
podman image rm my-go-api-helm:1.0
```
