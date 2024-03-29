output "manifests" {
  value = {
    "example" = data.k8s_pgv2_percona_com_percona_pg_backup_v2_manifest.example.yaml
  }
}
