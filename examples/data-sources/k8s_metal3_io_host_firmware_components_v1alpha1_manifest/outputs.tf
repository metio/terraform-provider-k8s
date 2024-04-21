output "manifests" {
  value = {
    "example" = data.k8s_metal3_io_host_firmware_components_v1alpha1_manifest.example.yaml
  }
}
