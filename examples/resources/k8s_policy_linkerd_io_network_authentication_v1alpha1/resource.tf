resource "k8s_policy_linkerd_io_network_authentication_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    networks = []
  }
}
