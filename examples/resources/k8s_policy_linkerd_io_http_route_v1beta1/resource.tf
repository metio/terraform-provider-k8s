resource "k8s_policy_linkerd_io_http_route_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {}
}
