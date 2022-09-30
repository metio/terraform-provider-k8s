resource "k8s_couchbase_com_couchbase_bucket_v2" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
}
