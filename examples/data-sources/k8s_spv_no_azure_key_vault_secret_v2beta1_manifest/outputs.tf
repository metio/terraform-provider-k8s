output "manifests" {
  value = {
    "example" = data.k8s_spv_no_azure_key_vault_secret_v2beta1_manifest.example.yaml
  }
}
