output "resources" {
  value = {
    "minimal" = k8s_monitoring_coreos_com_prometheus_v1.minimal.yaml
  }
}
