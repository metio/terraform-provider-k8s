data "k8s_cilium_io_cilium_clusterwide_envoy_config_v2_manifest" "example" {
  metadata = {
    name = "some-name"
  }
}
