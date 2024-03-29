output "manifests" {
  value = {
    "example" = data.k8s_kuadrant_io_kuadrant_v1beta1_manifest.example.yaml
  }
}
