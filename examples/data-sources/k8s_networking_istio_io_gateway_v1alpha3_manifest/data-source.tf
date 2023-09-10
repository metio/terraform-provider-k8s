data "k8s_networking_istio_io_gateway_v1alpha3_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
