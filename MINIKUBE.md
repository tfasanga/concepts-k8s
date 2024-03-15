# Minikube on MacOS

Run minikube with socket_vmnet network:

```shell
brew install qemu
brew install socket_vmnet
brew tap homebrew/services
HOMEBREW=$(which brew) && sudo ${HOMEBREW} services start socket_vmnet
```

If minikube was running before then stop & delete it:

```shell
minikube stop
minikube delete
```

```shell
minikube start --network socket_vmnet
```

or

```shell
minikube start --driver qemu --network socket_vmnet
```

Switch kubectl context to the "minikube" cluster:

```shell
kubectx minikube
```

# Build and Load docker image to Minikube

Build and Save Docker image:

```shell
podman build -t my-go-api-kustomize:1.0 .
podman save my-go-api-kustomize:1.0 -o my-go-api-kustomize-image.tar
```

```shell
minikube image load my-go-api-kustomize-image.tar
```

## Remove image from minikube

```shell
minikube image rm docker.io/localhost/my-go-api-kustomize:1.0
```

