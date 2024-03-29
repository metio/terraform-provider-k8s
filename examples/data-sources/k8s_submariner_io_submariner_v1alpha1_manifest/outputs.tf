output "manifests" {
  value = {
    "example" = data.k8s_submariner_io_submariner_v1alpha1_manifest.example.yaml
  }
}
