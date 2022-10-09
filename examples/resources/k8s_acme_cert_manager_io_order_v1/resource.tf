resource "k8s_acme_cert_manager_io_order_v1" "minimal" {
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
    request = "c29tZS1yZXF1ZXN0"
  }
}
