output "manifests" {
  value = {
    "example" = data.k8s_generators_external_secrets_io_sts_session_token_v1alpha1_manifest.example.yaml
  }
}
