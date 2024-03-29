data "k8s_datadoghq_com_datadog_agent_v2alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
