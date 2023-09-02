data "k8s_gateway_networking_k8s_io_tcp_route_v1alpha2_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
