output "resources" {
  value = {
    "minimal" = k8s_getambassador_io_consul_resolver_v2.minimal.yaml
  }
}
