resource "k8s_couchbase_com_couchbase_cluster_v2" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
  spec = {
    image = "some-image"
    security = {
      admin_secret = "some-secret"
    }
    servers = []
  }
}
