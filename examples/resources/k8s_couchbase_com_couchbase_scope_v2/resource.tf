resource "k8s_couchbase_com_couchbase_scope_v2" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
