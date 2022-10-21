output "resources" {
  value = {
    "minimal" = k8s_getambassador_io_auth_service_v2.minimal.yaml
  }
}
