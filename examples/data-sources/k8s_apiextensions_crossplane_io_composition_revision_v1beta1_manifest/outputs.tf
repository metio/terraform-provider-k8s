output "manifests" {
  value = {
    "example" = data.k8s_apiextensions_crossplane_io_composition_revision_v1beta1_manifest.example.yaml
  }
}
