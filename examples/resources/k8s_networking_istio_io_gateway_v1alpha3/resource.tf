resource "k8s_networking_istio_io_gateway_v1alpha3" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
