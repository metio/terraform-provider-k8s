output "manifests" {
  value = {
    "example" = data.k8s_admissionregistration_k8s_io_mutating_webhook_configuration_v1_manifest.example.yaml
  }
}
