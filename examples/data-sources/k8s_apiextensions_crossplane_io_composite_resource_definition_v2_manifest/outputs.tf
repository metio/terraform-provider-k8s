output "manifests" {
  value = {
    "example" = data.k8s_apiextensions_crossplane_io_composite_resource_definition_v2_manifest.example.yaml
  }
}
