output "manifests" {
  value = {
    "example" = data.k8s_psmdb_percona_com_percona_server_mongo_db_backup_v1_manifest.example.yaml
  }
}
