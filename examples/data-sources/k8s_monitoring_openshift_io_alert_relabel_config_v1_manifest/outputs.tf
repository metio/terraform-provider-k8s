output "manifests" {
  value = {
    "example" = data.k8s_monitoring_openshift_io_alert_relabel_config_v1_manifest.example.yaml
  }
}
