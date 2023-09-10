resource "k8s_cilium_io_cilium_node_config_v2alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
