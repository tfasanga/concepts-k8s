# Build

## Build and Load image in Kubernetes

```shell
podman build -t $(minikube ip):5000/my-go-api-goclient:latest .
podman push $(minikube ip):5000/my-go-api-goclient:latest
```

# Install

## Install in Kubernetes

```shell
kubectl apply -k kustomize/overlays/dev_registry
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
kubectl delete -k kustomize/overlays/dev_registry
```

## Remove image from Podman

```shell
podman image rm my-go-api-goclient:latest
```

# Build Manifest

```shell
kustomize build kustomize/overlays/dev_registry > dev_registry_manifest.yaml
```
