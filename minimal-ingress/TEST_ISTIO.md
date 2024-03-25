# Install Istio

[ISTIO_GATEWAY.md](ISTIO_GATEWAY)

# Install

## Install in Kubernetes

```shell
kubectl apply -k kustomize/istio
```

# Run

Ensure minikube tunnel is running in another terminal:

```shell
minikube tunnel
```

## Run test application via Ingress Gateway

```shell
export INGRESS_HOST=$(kubectl get gtw minimal-gateway -o jsonpath='{.status.addresses[0].value}')
export INGRESS_PORT=$(kubectl get gtw minimal-gateway -o jsonpath='{.spec.listeners[?(@.name=="http")].port}')
echo "http://$INGRESS_HOST:$INGRESS_PORT/minimal"
```

```shell
curl -v "$INGRESS_HOST:$INGRESS_PORT/minimal"
```

# Remove

## Uninstall in Kubernetes

```shell
kubectl delete -k kustomize/istio
```

```shell
kubectl get all
```

## Uninstall Istio Ingress Gateway in Kubernetes using Helm

[ISTIO_GATEWAY.md](ISTIO_GATEWAY)
