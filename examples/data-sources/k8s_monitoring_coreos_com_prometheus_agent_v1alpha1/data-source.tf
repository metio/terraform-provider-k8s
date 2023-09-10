data "k8s_monitoring_coreos_com_prometheus_agent_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
