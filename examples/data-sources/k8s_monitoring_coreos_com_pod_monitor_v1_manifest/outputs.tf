output "manifests" {
  value = {
    "example" = data.k8s_monitoring_coreos_com_pod_monitor_v1_manifest.example.yaml
    "int_target_port" = data.k8s_monitoring_coreos_com_pod_monitor_v1_manifest.int_target_port.yaml
    "string_target_port" = data.k8s_monitoring_coreos_com_pod_monitor_v1_manifest.string_target_port.yaml
  }
}
