data "k8s_monitoring_openshift_io_alerting_rule_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
