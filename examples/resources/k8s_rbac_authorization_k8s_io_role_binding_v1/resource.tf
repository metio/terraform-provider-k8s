resource "k8s_rbac_authorization_k8s_io_role_binding_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
