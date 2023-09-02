output "manifests" {
  value = {
    "example" = data.k8s_longhorn_io_backing_image_data_source_v1beta1_manifest.example.yaml
  }
}
