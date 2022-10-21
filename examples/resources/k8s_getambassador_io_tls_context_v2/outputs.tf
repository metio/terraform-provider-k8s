output "resources" {
  value = {
    "minimal" = k8s_getambassador_io_tls_context_v2.minimal.yaml
  }
}
