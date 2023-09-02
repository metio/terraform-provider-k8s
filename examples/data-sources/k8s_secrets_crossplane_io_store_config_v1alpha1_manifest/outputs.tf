output "manifests" {
  value = {
    "example" = data.k8s_secrets_crossplane_io_store_config_v1alpha1_manifest.example.yaml
  }
}
