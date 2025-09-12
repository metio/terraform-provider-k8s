data "k8s_grafana_integreatly_org_grafana_alert_rule_group_v1beta1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
