output "resources" {
  value = {
    "minimal" = k8s_service_account_v1.minimal.yaml
    "example" = k8s_service_account_v1.example.yaml
  }
}
