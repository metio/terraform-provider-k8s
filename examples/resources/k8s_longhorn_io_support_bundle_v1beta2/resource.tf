resource "k8s_longhorn_io_support_bundle_v1beta2" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
