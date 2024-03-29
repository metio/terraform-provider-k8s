output "manifests" {
  value = {
    "example" = data.k8s_secrets_hashicorp_com_vault_dynamic_secret_v1beta1_manifest.example.yaml
  }
}
