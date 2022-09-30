output "resources" {
  value = {
    "minimal" = k8s_monitoring_coreos_com_probe_v1.minimal.yaml
  }
}
