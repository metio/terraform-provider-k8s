output "resources" {
  value = {
    "minimal" = k8s_monitoring_coreos_com_thanos_ruler_v1.minimal.yaml
  }
}
