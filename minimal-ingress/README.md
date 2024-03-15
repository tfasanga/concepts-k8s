Uses: Podman, Minikube, Kubectl, Kustomize, Nginx Ingress Controller


# Download Nginx

```shell
mkdir nginx
wget https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v1.10.0/deploy/static/provider/baremetal/deploy.yaml -O nginx/deploy.yaml
```

# Prerequisite: Install Nginx Ingress Controller in Kubernetes

```shell
kubectl apply -f nginx/deploy.yaml
```

# Base

[TEST_BASE.md](TEST_BASE.md)

# Dev Overlay

[TEST_OVERLAY_DEV.md](TEST_OVERLAY_DEV.md)

# Remove Nginx Ingress Controller from Kubernetes

```shell
kubectl delete -f nginx/deploy.yaml
```
