resource "k8s_couchbase_com_couchbase_collection_group_v2" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
  spec = {
    names = ["name-1", "name-2"]
  }
}
