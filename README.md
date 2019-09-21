Kubernetes Config from CLI Flags
================================

A simple package that help parse CLI flags to generate client-go `*rest.Config`.

Three common CLI flags:

* `-kubeconfig`: path to the kubeconfig file to use (default to `~/.kube/config`)
* `-context`: the context within the kubeconfig file to use
* `-in-cluster`: use in-cluster configuration

