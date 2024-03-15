# Kind

```shell
brew install kind
kind create cluster
```

Switch kubectl context to the "kind-kind" cluster:

```shell
kubectx kind-kind
```

# Load docker image to Kind

Build and Save Docker image:

```shell
podman build -t my-go-api-kustomize:1.0 .
podman save my-go-api-kustomize:1.0 -o my-go-api-kustomize-image.tar
```

```shell
kind load image-archive my-go-api-kustomize-image.tar
```
