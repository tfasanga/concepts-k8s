Create

```shell
git clone https://github.com/tfasanga/learning-k8s.git
cd learning-k8s
mkdir minimal-go-manual
cd minimal-go-manual
go mod init fasanga/learn/k8s/manual
```

```shell
git clone https://github.com/tfasanga/learning-k8s.git
cd learning-k8s/minimal-go-manual
```

```shell
minikube start
```

Build Docker image:

```shell
podman build -t my-go-api-helm:1.0 .
```

Test run in Podman:

```shell
podman run --rm -p 8080:8080 --name my-go-api-helm my-go-api-helm:1.0 
```

Load docker image to minikube:

```shell
podman save my-go-api-helm:1.0 -o my-go-api-helm-image.tar
minikube image load my-go-api-helm-image.tar
```

Install Helm chart:

```shell
helm install my-go-api-chart helm --values helm/values.yaml
```

```shell
kubectl logs -l app=go-api-label -f
```

```shell
export POD_NAME=$(kubectl get pods -l "app=go-api-label" -o jsonpath="{.items[0].metadata.name}")
echo "$POD_NAME"
kubectl port-forward $POD_NAME 9000:9000
```

```shell
helm install my-go-api-chart helm --values env/dev-values.yaml
```

```shell
kubectl get pods
```

```shell
kubectl describe service
kubectl get service
```

Uninstall Helm chart:

```shell
helm uninstall my-go-api-chart
```

Remove image from minikube:

```shell
minikube image rm docker.io/localhost/my-go-api-helm:1.0
```

Remove image from podman:

```shell
podman image rm my-go-api-helm:1.0
```

