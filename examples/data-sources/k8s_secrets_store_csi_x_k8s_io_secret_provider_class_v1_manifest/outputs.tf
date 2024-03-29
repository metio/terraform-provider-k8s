output "manifests" {
  value = {
    "example" = data.k8s_secrets_store_csi_x_k8s_io_secret_provider_class_v1_manifest.example.yaml
  }
}
