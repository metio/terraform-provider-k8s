output "resources" {
  value = {
    "minimal" = k8s_cilium_io_cilium_node_v2.minimal.yaml
  }
}
