# Build

## Build and Load image in Kubernetes

```shell
podman build -t my-go-api-goclient:latest .
podman save my-go-api-goclient:latest -o my-go-api-goclient-image.tar
minikube image load my-go-api-goclient-image.tar
```

# Install

## Install in Kubernetes

```shell
kubectl apply -k kustomize/overlays/dev
```

```shell
kubectl get pods
```

```shell
kubectl describe pods
```

```shell
kubectl logs -f $(kubectl get pods -l "app=go-api-label" -o jsonpath="{.items[0].metadata.name}")
```

# Run

Note: This dev overlay uses NodePort service.
Get the service URL: `minikube service go-api-service --url`

```shell
curl $(minikube service go-api-service --url)
```

# Remove

## Uninstall in Kubernetes

```shell
kubectl delete -k kustomize/overlays/dev
```

## Remove image from Kubernetes

```shell
minikube image rm docker.io/localhost/my-go-api-goclient:latest
```

## Remove image from Podman

```shell
podman image rm my-go-api-goclient:1.0
```

# Build Manifest

```shell
goclient build kustomize/overlays/dev > dev_manifest.yaml
```
