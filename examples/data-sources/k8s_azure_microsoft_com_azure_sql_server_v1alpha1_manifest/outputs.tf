output "manifests" {
  value = {
    "example" = data.k8s_azure_microsoft_com_azure_sql_server_v1alpha1_manifest.example.yaml
  }
}
