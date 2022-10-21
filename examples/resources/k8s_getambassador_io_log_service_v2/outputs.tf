output "resources" {
  value = {
    "minimal" = k8s_getambassador_io_log_service_v2.minimal.yaml
  }
}
