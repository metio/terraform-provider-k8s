output "manifests" {
  value = {
    "example" = data.k8s_monitoring_coreos_com_prometheus_agent_v1alpha1_manifest.example.yaml
  }
}
