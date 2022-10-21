output "resources" {
  value = {
    "minimal" = k8s_getambassador_io_rate_limit_service_v2.minimal.yaml
  }
}
