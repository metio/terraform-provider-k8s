output "manifests" {
  value = {
    "example" = data.k8s_getambassador_io_rate_limit_service_v3alpha1_manifest.example.yaml
  }
}
