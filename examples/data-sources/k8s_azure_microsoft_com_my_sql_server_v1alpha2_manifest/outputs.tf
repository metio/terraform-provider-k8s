output "manifests" {
  value = {
    "example" = data.k8s_azure_microsoft_com_my_sql_server_v1alpha2_manifest.example.yaml
  }
}
