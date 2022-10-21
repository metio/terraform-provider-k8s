resource "k8s_cilium_io_cilium_node_v2" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {}
}
