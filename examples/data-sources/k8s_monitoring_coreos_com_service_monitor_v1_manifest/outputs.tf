output "manifests" {
  value = {
    "example" = data.k8s_monitoring_coreos_com_service_monitor_v1_manifest.example.yaml
  }
}
