output "resources" {
  value = {
    "minimal" = k8s_getambassador_io_kubernetes_endpoint_resolver_v2.minimal.yaml
  }
}
