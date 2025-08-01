output "manifests" {
  value = {
    "example" = data.k8s_opensearch_opster_io_opensearch_action_group_v1_manifest.example.yaml
  }
}
