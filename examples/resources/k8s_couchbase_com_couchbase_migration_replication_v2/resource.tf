resource "k8s_couchbase_com_couchbase_migration_replication_v2" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
  spec = {
    bucket        = "some-bucket"
    remote_bucket = "some-remote-bucket"
  }
  migration_mapping = {
    mappings = []
  }
}
