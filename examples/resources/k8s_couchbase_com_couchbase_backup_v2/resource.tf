resource "k8s_couchbase_com_couchbase_backup_v2" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}