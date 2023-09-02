output "manifests" {
  value = {
    "example" = data.k8s_longhorn_io_setting_v1beta1_manifest.example.yaml
  }
}
