output "resources" {
  value = {
    "minimal" = k8s_persistent_volume_v1.minimal.yaml
    "example" = k8s_persistent_volume_v1.example.yaml
  }
}
