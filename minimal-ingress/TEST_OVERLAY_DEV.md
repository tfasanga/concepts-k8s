# Build

## Build and Load image in Kubernetes

```shell
podman build -t my-go-api-ingress:1.0 .
podman save my-go-api-ingress:1.0 -o my-go-api-ingress-image.tar
minikube image load my-go-api-ingress-image.tar
```

# Install

## Install in Kubernetes

```shell
kubectl apply -k kustomize/overlays/dev
```

# Run

Note: This dev overlay uses NodePort service.
Get the service URL:

```shell
minikube service go-api-service --url
```

## Run test application via NodePort

```shell
curl $(minikube service go-api-service --url)
```

## Run test application via Ingress

```shell
curl "$(minikube -n ingress-nginx service ingress-nginx-controller --url | head -1)/minimal"
```

# Remove

## Uninstall in Kubernetes

```shell
kubectl delete -k kustomize/overlays/dev
```

## Remove image from Kubernetes

```shell
minikube image rm docker.io/localhost/my-go-api-ingress:1.0
```

## Remove image from Podman

```shell
podman image rm my-go-api-ingress:1.0
```

# Build Manifest

```shell
kustomize build kustomize/overlays/dev > dev_manifest.yaml
```

# Minikube SSH

```shell
minikube ssh
```

inside minikube ssh:

```shell
curl http://127.0.0.1:30090
```
