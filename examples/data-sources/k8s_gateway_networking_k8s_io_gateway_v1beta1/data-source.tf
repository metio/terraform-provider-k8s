data "k8s_gateway_networking_k8s_io_gateway_v1beta1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
