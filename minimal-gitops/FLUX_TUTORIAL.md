Uses: Podman, Minikube, Kubectl, Kustomize, Flux (GitOps)

https://fluxcd.io/flux/get-started/

Install Flux CLI:

```shell
brew install fluxcd/tap/flux
```
Assuming your token is saved in file `~/.ssh/github`:

```
export GITHUB_TOKEN=<your-token>
export GITHUB_USER=<your-username>
```

```shell
flux check --pre
```

Run the bootstrap command:

```shell
flux bootstrap github \
  --owner=$GITHUB_USER \
  --repository=fleet-infra \
  --branch=main \
  --path=./clusters/my-cluster \
  --personal
```

installing components in "flux-system" namespace

The bootstrap will create a new repo in github.com.
Clone the repo:

```shell
git clone https://github.com/$GITHUB_USER/fleet-infra
cd fleet-infra
```

List all k8s resources in namespace flux-system:

```shell
kubectl get all --namespace flux-system
```

```
NAME                                           READY   STATUS    RESTARTS   AGE
pod/helm-controller-5d8d5fc6fd-smfmr           1/1     Running   0          3m28s
pod/kustomize-controller-7b7b47f459-5j7xd      1/1     Running   0          3m28s
pod/notification-controller-5bb6647999-f6cm6   1/1     Running   0          3m28s
pod/source-controller-7667765cd7-2jqb4         1/1     Running   0          3m28s

NAME                              TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
service/notification-controller   ClusterIP   10.97.255.137   <none>        80/TCP    3m28s
service/source-controller         ClusterIP   10.101.44.67    <none>        80/TCP    3m28s
service/webhook-receiver          ClusterIP   10.102.127.8    <none>        80/TCP    3m28s

NAME                                      READY   UP-TO-DATE   AVAILABLE   AGE
deployment.apps/helm-controller           1/1     1            1           3m28s
deployment.apps/kustomize-controller      1/1     1            1           3m28s
deployment.apps/notification-controller   1/1     1            1           3m28s
deployment.apps/source-controller         1/1     1            1           3m28s

NAME                                                 DESIRED   CURRENT   READY   AGE
replicaset.apps/helm-controller-5d8d5fc6fd           1         1         1       3m28s
replicaset.apps/kustomize-controller-7b7b47f459      1         1         1       3m28s
replicaset.apps/notification-controller-5bb6647999   1         1         1       3m28s
replicaset.apps/source-controller-7667765cd7         1         1         1       3m28s
```

