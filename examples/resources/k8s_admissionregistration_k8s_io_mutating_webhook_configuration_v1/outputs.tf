output "resources" {
  value = {
    "minimal" = k8s_admissionregistration_k8s_io_mutating_webhook_configuration_v1.minimal.yaml
    "example" = k8s_admissionregistration_k8s_io_mutating_webhook_configuration_v1.example.yaml
  }
}
