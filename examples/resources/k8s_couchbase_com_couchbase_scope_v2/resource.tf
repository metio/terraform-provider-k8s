resource "k8s_couchbase_com_couchbase_scope_v2" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
}
