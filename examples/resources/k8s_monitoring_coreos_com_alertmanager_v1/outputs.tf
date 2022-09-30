output "resources" {
  value = {
    "minimal" = k8s_monitoring_coreos_com_alertmanager_v1.minimal.yaml
  }
}
