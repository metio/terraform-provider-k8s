output "manifests" {
  value = {
    "example" = data.k8s_azure_microsoft_com_key_vault_v1alpha1_manifest.example.yaml
  }
}
