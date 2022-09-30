output "resources" {
  value = {
    "minimal" = k8s_flagger_app_alert_provider_v1beta1.big.yaml
    "small"   = k8s_flagger_app_alert_provider_v1beta1.small.yaml
  }
}
