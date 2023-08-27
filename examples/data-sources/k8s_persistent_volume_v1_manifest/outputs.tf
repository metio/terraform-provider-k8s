output "manifests" {
  value = {
    "example" = data.k8s_persistent_volume_v1_manifest.example.yaml
  }
}
