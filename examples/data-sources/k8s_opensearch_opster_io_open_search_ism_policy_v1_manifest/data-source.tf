data "k8s_opensearch_opster_io_open_search_ism_policy_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
