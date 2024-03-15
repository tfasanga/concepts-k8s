# Go

```shell
go get github.com/anthdm/hollywood
go get k8s.io/apimachinery
go get k8s.io/api
go get k8s.io/client-go
go get k8s.io/klog/v2
go get github.com/spf13/viper
```

# In Cluster

If you have RBAC enabled on your cluster,
run the following snippet to create role binding
which will grant the default service account view permissions.

```shell
kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=default:default
```

If your application runs in a Pod in the cluster:
https://github.com/kubernetes/client-go/blob/v0.29.2/examples/in-cluster-client-configuration

# In Cluster without Docker registry

[TEST_IN_CLUSTER.md](TEST_IN_CLUSTER.md)

# In Cluster with Docker registry

[TEST_IN_CLUSTER_REGISTRY.md](TEST_IN_CLUSTER_REGISTRY.md)

# Out of Cluster (on host)

[TEST_OUT_OF_CLUSTER.md](TEST_OUT_OF_CLUSTER.md)

