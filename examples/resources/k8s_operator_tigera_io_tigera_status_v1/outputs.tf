output "resources" {
  value = {
    "minimal" = k8s_operator_tigera_io_tigera_status_v1.minimal.yaml
  }
}
