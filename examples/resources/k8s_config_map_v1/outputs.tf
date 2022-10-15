output "resources" {
  value = {
    "minimal" = k8s_config_map_v1.minimal.yaml
    "example" = k8s_config_map_v1.example.yaml
  }
}
