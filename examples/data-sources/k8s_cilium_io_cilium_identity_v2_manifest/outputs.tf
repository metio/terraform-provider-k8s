output "manifests" {
  value = {
    "example" = data.k8s_cilium_io_cilium_identity_v2_manifest.example.yaml
  }
}
