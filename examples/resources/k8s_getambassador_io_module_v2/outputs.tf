output "resources" {
  value = {
    "minimal" = k8s_getambassador_io_module_v2.minimal.yaml
  }
}
