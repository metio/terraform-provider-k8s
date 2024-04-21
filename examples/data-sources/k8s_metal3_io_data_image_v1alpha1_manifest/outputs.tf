output "manifests" {
  value = {
    "example" = data.k8s_metal3_io_data_image_v1alpha1_manifest.example.yaml
  }
}
