output "manifests" {
  value = {
    "example" = data.k8s_longhorn_io_backing_image_v1beta1_manifest.example.yaml
  }
}
