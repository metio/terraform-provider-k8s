resource "k8s_policy_linkerd_io_server_authorization_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    client = {}
    server = {}
  }
}
