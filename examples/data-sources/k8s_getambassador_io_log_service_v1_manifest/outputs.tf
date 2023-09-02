output "manifests" {
  value = {
    "example" = data.k8s_getambassador_io_log_service_v1_manifest.example.yaml
  }
}
