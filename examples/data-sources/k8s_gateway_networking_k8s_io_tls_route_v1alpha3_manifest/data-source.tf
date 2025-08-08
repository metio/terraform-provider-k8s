data "k8s_gateway_networking_k8s_io_tls_route_v1alpha3_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
