output "manifests" {
  value = {
    "example" = data.k8s_opensearch_opster_io_opensearch_index_template_v1_manifest.example.yaml
  }
}
