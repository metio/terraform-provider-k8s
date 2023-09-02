output "manifests" {
  value = {
    "example" = data.k8s_image_toolkit_fluxcd_io_image_repository_v1beta2_manifest.example.yaml
  }
}
