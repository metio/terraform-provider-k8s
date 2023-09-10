data "k8s_couchbase_com_couchbase_role_binding_v2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    subjects = []
    role_ref = {
      kind = "CouchbaseGroup"
      name = "some-name"
    }
  }
}
