data "k8s_gateway_networking_k8s_io_http_route_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
