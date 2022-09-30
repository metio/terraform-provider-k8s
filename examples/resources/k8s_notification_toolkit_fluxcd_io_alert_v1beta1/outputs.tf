output "resources" {
  value = {
    "big"   = k8s_notification_toolkit_fluxcd_io_alert_v1beta1.big.yaml
    "small" = k8s_notification_toolkit_fluxcd_io_alert_v1beta1.small.yaml
  }
}
