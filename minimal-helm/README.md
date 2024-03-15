Uses: Podman, Minikube, Kubectl, Helm

# Base

Uses: ClusterIP, Kubectl Port-Forwarding

[TEST_BASE.md](TEST_BASE.md)

# Dev Overlay

Uses: NodePort

[TEST_OVERLAY_DEV.md](TEST_OVERLAY_DEV.md)

# Notes

Show logs:

```shell
kubectl logs -l app=go-api-label -f
```

Show pods:

```shell
kubectl get pods
```

Show service:

```shell
kubectl describe service go-api-service
kubectl get service
```

