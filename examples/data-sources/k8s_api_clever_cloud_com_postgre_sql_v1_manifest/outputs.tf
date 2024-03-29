output "manifests" {
  value = {
    "example" = data.k8s_api_clever_cloud_com_postgre_sql_v1_manifest.example.yaml
  }
}
