output "manifests" {
  value = {
    "example" = data.k8s_projectcontour_io_contour_configuration_v1alpha1_manifest.example.yaml
  }
}
