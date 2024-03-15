Using Minikube and Podman

```shell
minikube start
```

```shell
podman build -t my-go-api:1.0 .
```

```shell
podman run -p 8080:8080 my-go-api:1.0
```

Load docker image to minikube:

```shell
podman save my-go-api:1.0 -o my-go-api-image.tar
minikube image load my-go-api-image.tar
minikube image ls
```

Remember to turn off the `imagePullPolicy:Always`
(use `imagePullPolicy:IfNotPresent` or `imagePullPolicy:Never`).

Manually deploy to Kubernetes:

```shell
kubectl apply -f config/default/deployment.yaml
```

```shell
kubectl get pods
```

Remove deployment:

```shell
kubectl delete -f config/default/deployment.yaml
```

Remove image:

```shell
minikube image ls
minikube image rm docker.io/localhost/my-go-api:1.0
```

```shell
podman image ls
podman image rm my-go-api:1.0
```

