output "manifests" {
  value = {
    "example" = data.k8s_cilium_io_cilium_clusterwide_network_policy_v2_manifest.example.yaml
  }
}
