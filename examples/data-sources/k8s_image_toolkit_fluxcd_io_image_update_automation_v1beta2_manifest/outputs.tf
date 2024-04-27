output "manifests" {
  value = {
    "example" = data.k8s_image_toolkit_fluxcd_io_image_update_automation_v1beta2_manifest.example.yaml
  }
}
