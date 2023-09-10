data "k8s_sematext_com_sematext_agent_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
