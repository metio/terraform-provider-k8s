output "manifests" {
  value = {
    "example" = data.k8s_projectcontour_io_extension_service_v1alpha1_manifest.example.yaml
  }
}
