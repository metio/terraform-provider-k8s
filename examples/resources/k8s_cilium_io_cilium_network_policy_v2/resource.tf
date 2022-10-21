resource "k8s_cilium_io_cilium_network_policy_v2" "minimal" {
  metadata = {
    name = "test"
  }
}
