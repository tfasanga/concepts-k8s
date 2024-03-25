# How to write Ingress Gateways resources

Ingress Gateways

https://istio.io/latest/docs/tasks/traffic-management/ingress/ingress-control/

Note that the Kubernetes Gateway API CRDs do not come installed by default on most Kubernetes clusters,
so make sure they are installed before using the Gateway API:

```shell
kubectl kustomize "github.com/kubernetes-sigs/gateway-api/config/crd?ref=444631bfe06f3bcca5d0eadf1857eac1d369421d" > manifests/gateway/crd.yaml
```

```shell
kubectl apply -f manifests/gateway/crd.yaml
```

If you are going to use the Gateway API instructions,
you can install Istio using the minimal profile because you will not need the istio-ingressgateway which is otherwise installed by default.


Gateway resource:

[minimal-ingress/kustomize/istio/gateway.yaml](minimal-ingress/kustomize/istio/gateway.yaml)

HTTP Route resource:

[minimal-ingress/kustomize/istio/http-route.yaml](minimal-ingress/kustomize/istio/http-route.yaml)
