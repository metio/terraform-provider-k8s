output "manifests" {
  value = {
    "example" = data.k8s_apiextensions_crossplane_io_composition_v1_manifest.example.yaml
  }
}
