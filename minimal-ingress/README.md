Uses: Podman, Minikube, Kubectl, Kustomize, Ingress Controller

## Build and Load image in Kubernetes

```shell
podman build -t my-go-api-ingress:1.0 .
podman save my-go-api-ingress:1.0 -o my-go-api-ingress-image.tar
minikube image load my-go-api-ingress-image.tar
```

## Run port forwarder

```shell
export POD_NAME=$(kubectl get pods -l "app=go-api-label" -o jsonpath="{.items[0].metadata.name}")
echo "$POD_NAME"
kubectl port-forward $POD_NAME 8080:8080
```

## Run test application

```shell
curl http://127.0.0.1:8080
```

# Istio

Install Istio: [ISTIO.md](ISTIO.md)

Write Gateway Resources: [ISTIO_GATEWAY.md](ISTIO_GATEWAY.md)

Run: [TEST_ISTIO](TEST_ISTIO.md)

# Nginx

[TEST_NGINX.md](TEST_NGINX.md)

