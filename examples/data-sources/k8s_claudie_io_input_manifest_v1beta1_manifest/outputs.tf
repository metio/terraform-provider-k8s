output "manifests" {
  value = {
    "example" = data.k8s_claudie_io_input_manifest_v1beta1_manifest.example.yaml
  }
}
