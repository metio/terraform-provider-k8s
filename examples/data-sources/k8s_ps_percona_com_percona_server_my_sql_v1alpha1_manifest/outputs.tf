output "manifests" {
  value = {
    "example" = data.k8s_ps_percona_com_percona_server_my_sql_v1alpha1_manifest.example.yaml
  }
}
