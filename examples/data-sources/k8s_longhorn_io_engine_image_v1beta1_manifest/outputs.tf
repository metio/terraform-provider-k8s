output "manifests" {
  value = {
    "example" = data.k8s_longhorn_io_engine_image_v1beta1_manifest.example.yaml
  }
}
