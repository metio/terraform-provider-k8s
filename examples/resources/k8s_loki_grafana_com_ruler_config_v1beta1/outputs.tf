output "resources" {
  value = {
    "minimal" = k8s_loki_grafana_com_ruler_config_v1beta1.minimal.yaml
  }
}
