output "resources" {
  value = {
    "minimal" = k8s_gateway_solo_io_virtual_service_v1.minimal.yaml
  }
}
