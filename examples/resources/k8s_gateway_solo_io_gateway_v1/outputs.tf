output "resources" {
  value = {
    "minimal" = k8s_gateway_solo_io_gateway_v1.minimal.yaml
  }
}
