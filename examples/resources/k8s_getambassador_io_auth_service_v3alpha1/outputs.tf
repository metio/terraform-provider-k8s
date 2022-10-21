output "resources" {
  value = {
    "minimal" = k8s_getambassador_io_auth_service_v3alpha1.minimal.yaml
  }
}
