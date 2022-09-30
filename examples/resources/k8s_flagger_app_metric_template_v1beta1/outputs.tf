output "resources" {
  value = {
    "minimal" = k8s_flagger_app_metric_template_v1beta1.minimal.yaml
  }
}
