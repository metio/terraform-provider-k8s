data "k8s_couchbase_com_couchbase_backup_restore_v2_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
