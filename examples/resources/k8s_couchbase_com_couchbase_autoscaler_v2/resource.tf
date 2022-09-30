resource "k8s_couchbase_com_couchbase_autoscaler_v2" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
  spec = {
    servers = "some-servers"
    size    = 17
  }
}
