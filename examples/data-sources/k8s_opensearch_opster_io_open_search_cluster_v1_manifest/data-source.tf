data "k8s_opensearch_opster_io_open_search_cluster_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
