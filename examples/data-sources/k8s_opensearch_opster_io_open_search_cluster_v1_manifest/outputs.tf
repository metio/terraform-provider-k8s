output "manifests" {
  value = {
    "example" = data.k8s_opensearch_opster_io_open_search_cluster_v1_manifest.example.yaml
  }
}
