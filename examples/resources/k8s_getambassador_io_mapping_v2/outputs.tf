output "resources" {
  value = {
    "minimal" = k8s_getambassador_io_mapping_v2.minimal.yaml
  }
}
