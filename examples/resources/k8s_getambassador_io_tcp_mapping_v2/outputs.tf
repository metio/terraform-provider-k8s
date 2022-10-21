output "resources" {
  value = {
    "minimal" = k8s_getambassador_io_tcp_mapping_v2.minimal.yaml
  }
}
