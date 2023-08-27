output "manifests" {
  value = {
    "example" = data.k8s_config_map_v1_manifest.example.yaml
  }
}
