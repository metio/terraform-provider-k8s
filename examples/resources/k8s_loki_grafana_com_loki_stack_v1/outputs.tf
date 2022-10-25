output "resources" {
  value = {
    "minimal" = k8s_loki_grafana_com_loki_stack_v1.minimal.yaml
  }
}
