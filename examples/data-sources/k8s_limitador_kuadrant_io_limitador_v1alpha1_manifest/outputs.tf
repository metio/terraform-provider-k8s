output "manifests" {
  value = {
    "example" = data.k8s_limitador_kuadrant_io_limitador_v1alpha1_manifest.example.yaml
  }
}
