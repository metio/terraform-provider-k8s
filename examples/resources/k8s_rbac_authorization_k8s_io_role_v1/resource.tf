resource "k8s_rbac_authorization_k8s_io_role_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_rbac_authorization_k8s_io_role_v1" "example" {
  metadata = {
    name = "test"
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
