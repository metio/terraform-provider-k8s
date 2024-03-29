output "manifests" {
  value = {
    "example" = data.k8s_operator_authorino_kuadrant_io_authorino_v1beta1_manifest.example.yaml
  }
}
