output "resources" {
  value = {
    "minimal" = k8s_getambassador_io_host_v2.minimal.yaml
  }
}
