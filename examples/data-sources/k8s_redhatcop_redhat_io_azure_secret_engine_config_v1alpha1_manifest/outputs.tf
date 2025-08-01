output "manifests" {
  value = {
    "example" = data.k8s_redhatcop_redhat_io_azure_secret_engine_config_v1alpha1_manifest.example.yaml
  }
}
