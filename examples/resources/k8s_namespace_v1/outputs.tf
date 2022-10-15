output "resources" {
  value = {
    "minimal" = k8s_namespace_v1.minimal.yaml
    "example" = k8s_namespace_v1.example.yaml
  }
}
