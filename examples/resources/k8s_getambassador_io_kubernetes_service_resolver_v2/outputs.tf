output "resources" {
  value = {
    "minimal" = k8s_getambassador_io_kubernetes_service_resolver_v2.minimal.yaml
  }
}
