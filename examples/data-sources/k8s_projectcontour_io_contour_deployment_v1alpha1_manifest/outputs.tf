output "manifests" {
  value = {
    "example" = data.k8s_projectcontour_io_contour_deployment_v1alpha1_manifest.example.yaml
  }
}
