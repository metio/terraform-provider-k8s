output "manifests" {
  value = {
    "example" = data.k8s_grafana_integreatly_org_grafana_alert_rule_group_v1beta1_manifest.example.yaml
  }
}
