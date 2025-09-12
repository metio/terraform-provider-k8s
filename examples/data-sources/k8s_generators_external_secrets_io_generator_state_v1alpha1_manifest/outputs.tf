output "manifests" {
  value = {
    "example" = data.k8s_generators_external_secrets_io_generator_state_v1alpha1_manifest.example.yaml
  }
}
