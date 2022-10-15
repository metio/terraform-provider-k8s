output "resources" {
  value = {
    "minimal" = k8s_secret_v1.minimal.yaml
    "example" = k8s_secret_v1.example.yaml
  }
}
