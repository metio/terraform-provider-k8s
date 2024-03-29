output "manifests" {
  value = {
    "example" = data.k8s_notification_toolkit_fluxcd_io_alert_v1beta3_manifest.example.yaml
  }
}
