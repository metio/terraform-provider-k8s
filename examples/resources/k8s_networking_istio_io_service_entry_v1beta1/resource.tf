resource "k8s_networking_istio_io_service_entry_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}
