resource "k8s_couchbase_com_couchbase_user_v2" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
  spec = {
    auth_domain = "some-domain"
  }
}
