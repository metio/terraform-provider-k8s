output "resources" {
  value = {
    "minimal" = k8s_loki_grafana_com_recording_rule_v1beta1.minimal.yaml
  }
}
