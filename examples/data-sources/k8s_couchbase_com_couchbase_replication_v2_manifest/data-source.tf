data "k8s_couchbase_com_couchbase_replication_v2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    bucket        = "some-bucket"
    remote_bucket = "some-remote-bucket"
  }
}
