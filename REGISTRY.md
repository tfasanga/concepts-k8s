# Local Docker Registry in Minikube on MacOS

```shell
minikube addons enable registry
```

Verify:

```shell
curl http://$(minikube ip):5000/v2/
```

returns `{}`

# Enable insecure registry in Podman

Determine Minikube IP:

```shell
minikube ip
```

returns `192.168.105.5`

*Note: On Mac OS Podman runs inside a VM.*

Enter Podman VM:

```shell
podman machine ssh
```

```
sudo vi /etc/containers/registries.conf
```

Add following section:

```
[[registry]]
location = "192.168.105.5:5000"
insecure = true
```

```
# Exit Podman VM
exit
```

Restart Podman VM:

```shell
podman machine stop
podman machine start
```


# Build and push Docker image to local registry

```shell
podman build -t $(minikube ip):5000/my-go-api-goclient:1.0 .
podman push $(minikube ip):5000/my-go-api-goclient:1.0
```

Image name in YAML will be `localhost:5000/my-go-api-goclient`.

Verify:

```shell
curl -X GET http://$(minikube ip):5000/v2/_catalog
```

returns:

```
{"repositories":["my-go-api-goclient"]}
```

```shell
curl -X GET http://$(minikube ip):5000/v2/my-go-api-goclient/tags/list
```

returns:

```
{"name":"my-go-api-goclient","tags":["1.0"]}
```
