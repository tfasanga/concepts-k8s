# Build

## Build and Load image in Kubernetes

```shell
podman build -t my-go-api-helm:1.0 .
podman save my-go-api-helm:1.0 -o my-go-api-helm-image.tar
minikube image load my-go-api-helm-image.tar
```

# Install

## Install in Kubernetes

```shell
helm install go-api helm --values env/dev-values.yaml
```

# Run

Note: This dev overlay uses NodePort service.
Get the service URL:

```shell
minikube service go-api-service --url
```

## Run test application

```shell
curl $(minikube service go-api-service --url)
```

# Remove

## Uninstall in Kubernetes

```shell
helm uninstall go-api
```

## Remove image from Kubernetes

```shell
minikube image rm docker.io/localhost/my-go-api-helm:1.0
```

## Remove image from Podman

```shell
podman image rm my-go-api-helm:1.0
```

# Minikube SSH

```shell
minikube ssh
```

inside minikube ssh:

```shell
curl http://127.0.0.1:30090
```
