data "k8s_opensearch_opster_io_opensearch_user_role_binding_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
