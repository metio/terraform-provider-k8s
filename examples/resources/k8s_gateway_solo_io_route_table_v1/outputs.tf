output "resources" {
  value = {
    "minimal" = k8s_gateway_solo_io_route_table_v1.minimal.yaml
  }
}
