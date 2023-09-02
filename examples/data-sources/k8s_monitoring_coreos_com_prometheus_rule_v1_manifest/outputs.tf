output "manifests" {
  value = {
    "example" = data.k8s_monitoring_coreos_com_prometheus_rule_v1_manifest.example.yaml
  }
}
