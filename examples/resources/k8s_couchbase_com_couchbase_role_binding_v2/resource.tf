resource "k8s_couchbase_com_couchbase_role_binding_v2" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
  spec = {
    subjects = []
    role_ref = {
      kind = "some-kind"
      name = "some-name"
    }
  }
}
