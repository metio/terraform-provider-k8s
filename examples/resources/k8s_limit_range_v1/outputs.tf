output "resources" {
  value = {
    "minimal" = k8s_limit_range_v1.minimal.yaml
    "example" = k8s_limit_range_v1.example.yaml
  }
}
