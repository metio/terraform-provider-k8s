data "k8s_opensearch_opster_io_opensearch_role_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
