output "manifests" {
  value = {
    "example" = data.k8s_canaries_flanksource_com_topology_v1_manifest.example.yaml
  }
}
