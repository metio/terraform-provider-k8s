output "resources" {
  value = {
    "minimal" = k8s_monitoring_coreos_com_alertmanager_config_v1alpha1.minimal.yaml
  }
}
