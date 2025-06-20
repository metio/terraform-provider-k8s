output "manifests" {
  value = {
    "example" = data.k8s_generators_external_secrets_io_password_v1alpha1_manifest.example.yaml
  }
}
