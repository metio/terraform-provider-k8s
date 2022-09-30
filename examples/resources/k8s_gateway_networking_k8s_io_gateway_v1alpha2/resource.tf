resource "k8s_gateway_networking_k8s_io_gateway_v1alpha2" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    gateway_class_name = "some-class"
    listeners          = []
  }
}
