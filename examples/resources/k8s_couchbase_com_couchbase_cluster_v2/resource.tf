resource "k8s_couchbase_com_couchbase_cluster_v2" "minimal" {
  metadata = {
    name      = "test"
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

resource "k8s_couchbase_com_couchbase_cluster_v2" "sha_image" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
  spec = {
    image = "hub.docker.com/image@sha256:45b23dee08af5e43a7fea6c4cf9c25ccf269ee113168c19722f87876677c5cb2"
    security = {
      admin_secret = "some-secret"
    }
    servers = []
  }
}
