# Build

## Build and Load image in Kubernetes

```shell
podman build -t my-go-api-goclient:1.0 .
podman save my-go-api-goclient:1.0 -o my-go-api-goclient-image.tar
minikube image load my-go-api-goclient-image.tar
```

# Install

## Install in Kubernetes

```shell
kubectl apply -k kustomize/base
```

```shell
kubectl get pods
```

```shell
kubectl logs -f $(kubectl get pods -l "app=go-api-label" -o jsonpath="{.items[0].metadata.name}")
```

# Run

## Run port forwarder

```shell
kubectl port-forward $(kubectl get pods -l "app=go-api-label" -o jsonpath="{.items[0].metadata.name}") 8080:8080
```

## Run test application

```shell
curl http://127.0.0.1:8080
```

# Remove

## Uninstall in Kubernetes

```shell
kubectl delete -k kustomize/base
```

## Remove image from Kubernetes

```shell
minikube image rm docker.io/localhost/my-go-api-goclient:1.0
```

## Remove image from Podman

```shell
podman image rm my-go-api-goclient:1.0
```

# Build Manifest

```shell
goclient build goclient/base > manifest.yaml
```
