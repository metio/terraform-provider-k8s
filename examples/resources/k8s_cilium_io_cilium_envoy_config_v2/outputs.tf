output "resources" {
  value = {
    "minimal" = k8s_cilium_io_cilium_envoy_config_v2.minimal.yaml
  }
}
