output "manifests" {
  value = {
    "example" = data.k8s_external_secrets_io_secret_store_v1alpha1_manifest.example.yaml
  }
}
