resource "k8s_cert_manager_io_certificate_v1" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
  spec = {
    issuer_ref = {
      group = "some-group"
      kind  = "some-kind"
      name  = "some-name"
    }
    secret_name = "some-secret"
  }
}
