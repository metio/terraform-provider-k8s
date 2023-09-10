data "k8s_cert_manager_io_certificate_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
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
