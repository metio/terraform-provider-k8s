resource "k8s_couchbase_com_couchbase_backup_v2" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
  spec = {
    strategy = "full_incremental"
  }
}
