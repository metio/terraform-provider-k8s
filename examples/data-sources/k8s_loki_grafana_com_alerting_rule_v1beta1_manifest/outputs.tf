output "manifests" {
  value = {
    "example" = data.k8s_loki_grafana_com_alerting_rule_v1beta1_manifest.example.yaml
  }
}
