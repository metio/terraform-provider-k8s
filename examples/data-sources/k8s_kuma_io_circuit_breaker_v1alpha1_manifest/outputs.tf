output "manifests" {
  value = {
    "example" = data.k8s_kuma_io_circuit_breaker_v1alpha1_manifest.example.yaml
  }
}
