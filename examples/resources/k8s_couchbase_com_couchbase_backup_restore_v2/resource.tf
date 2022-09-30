resource "k8s_couchbase_com_couchbase_backup_restore_v2" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
  spec = {
    backup = "some-backup"
    repo   = "some-repo"
  }
}
