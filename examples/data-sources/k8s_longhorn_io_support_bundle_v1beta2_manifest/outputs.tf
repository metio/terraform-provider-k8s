output "manifests" {
  value = {
    "example" = data.k8s_longhorn_io_support_bundle_v1beta2_manifest.example.yaml
  }
}
