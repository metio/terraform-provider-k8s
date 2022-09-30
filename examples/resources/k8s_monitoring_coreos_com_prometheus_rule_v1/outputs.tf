output "resources" {
  value = {
    "minimal" = k8s_monitoring_coreos_com_prometheus_rule_v1.minimal.yaml
  }
}
