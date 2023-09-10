data "k8s_monitoring_coreos_com_prometheus_rule_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {}
}
