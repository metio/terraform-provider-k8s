output "manifests" {
  value = {
    "example" = data.k8s_networking_k8s_io_network_policy_v1_manifest.example.yaml
  }
}
