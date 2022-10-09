resource "k8s_cert_manager_io_certificate_request_v1" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
  spec = {
    request = "c29tZS1yZXF1ZXN0"
    issuer_ref = {
      group = "some-group"
      kind  = "some-kind"
      name  = "some-name"
    }
  }
}
