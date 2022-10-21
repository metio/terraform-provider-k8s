output "resources" {
  value = {
    "minimal" = k8s_getambassador_io_tracing_service_v3alpha1.minimal.yaml
  }
}
