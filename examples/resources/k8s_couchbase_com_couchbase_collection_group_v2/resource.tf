resource "k8s_couchbase_com_couchbase_collection_group_v2" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
