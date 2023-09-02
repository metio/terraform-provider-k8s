output "manifests" {
  value = {
    "example" = data.k8s_monitoring_coreos_com_pod_monitor_v1_manifest.example.yaml
  }
}
