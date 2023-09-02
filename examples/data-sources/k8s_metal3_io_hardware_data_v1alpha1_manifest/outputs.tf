output "manifests" {
  value = {
    "example" = data.k8s_metal3_io_hardware_data_v1alpha1_manifest.example.yaml
  }
}
