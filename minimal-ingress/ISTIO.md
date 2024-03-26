# Prerequisite: Install Istio Ingress Gateway in Kubernetes

## Prerequisites for Minikube

https://istio.io/latest/docs/setup/platform-setup/minikube/

# Installing using helm

https://istio.io/latest/docs/setup/install/helm/

```shell
helm repo add istio https://istio-release.storage.googleapis.com/charts
helm repo update
```

```shell
kubectl create namespace istio-system
```

Install the Istio base chart which contains cluster-wide Custom Resource Definitions (CRDs) which must be installed prior to the deployment of the Istio control plane:

```shell
helm install istio-base istio/base -n istio-system --set defaultRevision=default
```

```shell
helm ls -n istio-system
```

Install the Istio discovery chart which deploys the istiod service:


```shell
helm install istiod istio/istiod -n istio-system --wait
```

```shell
helm ls -n istio-system
```

```shell
helm status istiod -n istio-system
```

```shell
kubectl get deployments -n istio-system --output wide
```

# How to write Ingress Gateways resources

Ingress Gateways

[ISTIO_GATEWAY.md](ISTIO_GATEWAY)


# Uninstalling using helm

```shell
kubectl get all -n istio-system
```

```shell
helm ls -n istio-system
```

```shell
helm delete istiod -n istio-system
```

```shell
helm delete istio-base -n istio-system
```

By design, deleting a chart via Helm doesn’t delete the installed Custom Resource Definitions (CRDs) installed via the chart.

```shell
kubectl delete namespace istio-system
```

