output "manifests" {
  value = {
    "example" = data.k8s_flagger_app_alert_provider_v1beta1_manifest.example.yaml
  }
}
