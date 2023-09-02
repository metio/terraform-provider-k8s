output "manifests" {
  value = {
    "example" = data.k8s_databases_schemahero_io_database_v1alpha4_manifest.example.yaml
  }
}
