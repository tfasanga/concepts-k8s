# Install Nginx

[NGINX.md](NGINX.md)

# Install

## Install in Kubernetes

```shell
kubectl apply -k kustomize/nginx
```

# Run

## Run test application via Ingress

```shell
curl "$(minikube -n ingress-nginx service ingress-nginx-controller --url | head -1)/minimal"
```

# Remove

## Uninstall in Kubernetes

```shell
kubectl delete -k kustomize/nginx
```
