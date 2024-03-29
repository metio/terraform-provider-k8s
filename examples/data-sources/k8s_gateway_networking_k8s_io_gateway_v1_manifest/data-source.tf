data "k8s_gateway_networking_k8s_io_gateway_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    gateway_class_name = "some-class"
    listeners          = []
  }
}
