resource "k8s_gateway_networking_k8s_io_gateway_class_v1alpha2" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    controller_name = "some-controller"
  }
}
