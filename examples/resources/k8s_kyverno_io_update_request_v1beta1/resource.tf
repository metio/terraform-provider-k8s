resource "k8s_kyverno_io_update_request_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}
