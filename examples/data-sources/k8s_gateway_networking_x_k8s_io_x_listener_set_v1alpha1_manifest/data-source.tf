data "k8s_gateway_networking_x_k8s_io_x_listener_set_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
