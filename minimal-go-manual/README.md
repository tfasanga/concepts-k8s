Create

```shell
git clone https://github.com/tfasanga/learning-k8s.git
cd learning-k8s
mkdir minimal-go-manual
cd minimal-go-manual
go mod init fasanga/learn/k8s/manual
```

```shell
minikube start
```

```shell
git clone https://github.com/tfasanga/learning-k8s.git
cd learning-k8s/minimal-go-manual
```

```shell
podman build -t my-go-api:1.0 .
#podman run -p 8080:8080 my-go-api:1.0

```
Load docker image to minikube:

```shell
podman save my-go-api:1.0 -o my-go-api-image.tar
minikube image load my-go-api-image.tar
minikube image ls
```

Remember to turn off the imagePullPolicy:Always
(use imagePullPolicy:IfNotPresent or imagePullPolicy:Never),

```shell
kubectl apply -f config/default/deployment.yaml
kubectl get pods
```

```shell
kubectl delete -f config/default/deployment.yaml
```

```shell
minikube image ls
minikube image rm docker.io/localhost/my-go-api:1.0
podman image rm my-go-api:1.0
```