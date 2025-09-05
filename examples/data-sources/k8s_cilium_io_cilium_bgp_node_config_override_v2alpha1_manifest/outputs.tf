output "manifests" {
  value = {
    "example" = data.k8s_cilium_io_cilium_bgp_node_config_override_v2alpha1_manifest.example.yaml
  }
}
