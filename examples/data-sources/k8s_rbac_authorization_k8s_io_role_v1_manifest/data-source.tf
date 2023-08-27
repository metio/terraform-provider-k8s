data "k8s_rbac_authorization_k8s_io_role_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  rules = [
    {
      api_groups     = [""]
      resources      = ["pods"]
      resource_names = ["foo"]
      verbs          = ["get", "list", "watch"]
    },
    {
      api_groups = ["apps"]
      resources  = ["deployments"]
      verbs      = ["get", "list"]
    },
  ]
}
