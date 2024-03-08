# Build

## Build and Load image in Kubernetes

```shell
podman build -t my-go-api-kustomize:1.0 .
podman save my-go-api-kustomize:1.0 -o my-go-api-kustomize-image.tar
minikube image load my-go-api-kustomize-image.tar
```

# Install

## Install in Kubernetes

```shell
kubectl apply -k kustomize/base
```

# Run

## Run test application via Ingress

```shell
curl "$(minikube -n ingress-nginx service ingress-nginx-controller --url | head -1)/minimal"
```

Get URL and open in browser: 

```shell
echo "$(minikube -n ingress-nginx service ingress-nginx-controller --url | head -1)/minimal"
```

# Remove

## Uninstall in Kubernetes

```shell
kubectl delete -k kustomize/base
```

## Remove image from Kubernetes

```shell
minikube image rm docker.io/localhost/my-go-api-kustomize:1.0
```

## Remove image from Podman

```shell
podman image rm my-go-api-kustomize:1.0
```

# Build Manifest

```shell
kustomize build kustomize/base > manifest.yaml
```
