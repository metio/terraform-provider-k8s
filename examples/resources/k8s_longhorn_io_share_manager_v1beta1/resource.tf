resource "k8s_longhorn_io_share_manager_v1beta1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
