resource "k8s_monitoring_coreos_com_prometheus_rule_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
