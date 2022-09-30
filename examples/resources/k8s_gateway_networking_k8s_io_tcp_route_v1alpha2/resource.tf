resource "k8s_gateway_networking_k8s_io_tcp_route_v1alpha2" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    rules = []
  }
}
