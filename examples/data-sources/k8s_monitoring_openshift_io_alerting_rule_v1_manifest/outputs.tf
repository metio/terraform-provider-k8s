output "manifests" {
  value = {
    "example" = data.k8s_monitoring_openshift_io_alerting_rule_v1_manifest.example.yaml
  }
}
