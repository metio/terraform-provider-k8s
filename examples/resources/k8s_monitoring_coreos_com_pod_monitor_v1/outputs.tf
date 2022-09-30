output "resources" {
  value = {
    "minimal" = k8s_monitoring_coreos_com_pod_monitor_v1.minimal.yaml
  }
}
