# Download Nginx Kubernetes Manifests

```shell
mkdir nginx
wget https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.10.0/deploy/static/provider/baremetal/deploy.yaml -O manifests/nginx/deploy.yaml
```

# Install Nginx Ingress Controller in Kubernetes

```shell
kubectl apply -f manifests/nginx/deploy.yaml
```

# Remove Nginx Ingress Controller from Kubernetes

```shell
kubectl delete -f manifests/nginx/deploy.yaml
```
