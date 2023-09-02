output "manifests" {
  value = {
    "example" = data.k8s_fossul_io_fossul_v1_manifest.example.yaml
  }
}
