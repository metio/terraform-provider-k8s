resource "k8s_networking_istio_io_gateway_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}
