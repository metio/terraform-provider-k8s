resource "k8s_rbac_authorization_k8s_io_cluster_role_binding_v1" "minimal" {
  metadata = {
    name = "test"
  }
  role_ref = {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = "cluster-admin"
  }
}

resource "k8s_rbac_authorization_k8s_io_cluster_role_binding_v1" "example" {
  metadata = {
    name = "test"
  }
  role_ref = {
    api_group = "rbac.authorization.k8s.io"
    kind      = "ClusterRole"
    name      = "cluster-admin"
  }
  subjects = [
    {
      kind      = "User"
      name      = "admin"
      api_group = "rbac.authorization.k8s.io"
    },
    {
      kind      = "ServiceAccount"
      name      = "default"
      namespace = "kube-system"
    },
    {
      kind      = "Group"
      name      = "system:masters"
      api_group = "rbac.authorization.k8s.io"
    },
  ]
}
