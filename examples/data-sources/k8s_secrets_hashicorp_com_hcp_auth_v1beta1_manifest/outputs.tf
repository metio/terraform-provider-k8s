output "manifests" {
  value = {
    "example" = data.k8s_secrets_hashicorp_com_hcp_auth_v1beta1_manifest.example.yaml
  }
}
