data "k8s_rbac_authorization_k8s_io_cluster_role_v1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  rules = [
    {
      api_groups = [""]
      resources  = ["namespaces", "pods"]
      verbs      = ["get", "list", "watch"]
    },
  ]
}
