output "manifests" {
  value = {
    "example" = data.k8s_schemas_schemahero_io_data_type_v1alpha4_manifest.example.yaml
  }
}
