resource "k8s_rbac_authorization_k8s_io_cluster_role_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_rbac_authorization_k8s_io_cluster_role_v1" "example" {
  metadata = {
    name = "test"
  }
  rules = [
    {
      api_groups = [""]
      resources  = ["namespaces", "pods"]
      verbs      = ["get", "list", "watch"]
    },
  ]
}
