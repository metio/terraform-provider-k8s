output "resources" {
  value = {
    "minimal" = k8s_getambassador_io_tracing_service_v2.minimal.yaml
  }
}
