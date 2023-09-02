output "manifests" {
  value = {
    "example" = data.k8s_operator_tigera_io_image_set_v1_manifest.example.yaml
  }
}
