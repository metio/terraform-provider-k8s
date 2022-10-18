output "resources" {
  value = {
    "minimal"  = k8s_couchbase_com_couchbase_backup_v2.minimal.yaml
    "example1" = k8s_couchbase_com_couchbase_backup_v2.example1.yaml
    "example2" = k8s_couchbase_com_couchbase_backup_v2.example2.yaml
  }
}
