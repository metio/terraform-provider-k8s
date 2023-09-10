resource "k8s_cert_manager_io_issuer_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
