data "k8s_couchbase_com_couchbase_cluster_v2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    image = "hub.docker.com/orga/image:version1.2.3"
    security = {
      admin_secret = "some-secret"
    }
    servers = []
  }
}
