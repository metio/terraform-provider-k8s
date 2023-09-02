output "manifests" {
  value = {
    "example" = data.k8s_fossul_io_restore_v1_manifest.example.yaml
  }
}
