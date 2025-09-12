output "manifests" {
  value = {
    "example" = data.k8s_opensearch_opster_io_open_search_ism_policy_v1_manifest.example.yaml
  }
}
