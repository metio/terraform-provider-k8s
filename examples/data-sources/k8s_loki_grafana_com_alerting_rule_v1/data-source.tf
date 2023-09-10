data "k8s_loki_grafana_com_alerting_rule_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
