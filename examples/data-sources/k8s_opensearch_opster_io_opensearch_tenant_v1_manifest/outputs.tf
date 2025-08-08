output "manifests" {
  value = {
    "example" = data.k8s_opensearch_opster_io_opensearch_tenant_v1_manifest.example.yaml
  }
}
