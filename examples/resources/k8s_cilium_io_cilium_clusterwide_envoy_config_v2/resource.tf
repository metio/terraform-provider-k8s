resource "k8s_cilium_io_cilium_clusterwide_envoy_config_v2" "minimal" {
  metadata = {
    name = "test"
  }
}
