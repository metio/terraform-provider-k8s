output "manifests" {
  value = {
    "example" = data.k8s_azure_microsoft_com_my_sqlaad_user_v1alpha2_manifest.example.yaml
  }
}
