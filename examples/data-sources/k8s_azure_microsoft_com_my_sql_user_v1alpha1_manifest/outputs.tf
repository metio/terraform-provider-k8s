output "manifests" {
  value = {
    "example" = data.k8s_azure_microsoft_com_my_sql_user_v1alpha1_manifest.example.yaml
  }
}
