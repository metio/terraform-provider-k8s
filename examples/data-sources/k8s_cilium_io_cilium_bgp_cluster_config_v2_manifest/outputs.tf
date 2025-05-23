output "manifests" {
  value = {
    "example" = data.k8s_cilium_io_cilium_bgp_cluster_config_v2_manifest.example.yaml
  }
}
