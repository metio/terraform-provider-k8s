output "manifests" {
  value = {
    "example" = data.k8s_netchecks_io_network_assertion_v1_manifest.example.yaml
  }
}
