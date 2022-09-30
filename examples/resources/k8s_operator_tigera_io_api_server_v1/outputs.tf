output "resources" {
  value = {
    "minimal" = k8s_operator_tigera_io_api_server_v1.minimal.yaml
  }
}
