output "manifests" {
  value = {
    "example" = data.k8s_opensearch_opster_io_opensearch_user_role_binding_v1_manifest.example.yaml
  }
}
