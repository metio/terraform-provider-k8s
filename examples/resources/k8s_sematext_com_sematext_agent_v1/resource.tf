resource "k8s_sematext_com_sematext_agent_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_sematext_com_sematext_agent_v1" "example" {
  metadata = {
    name = "basic-agent-deployment"
  }
  spec = {
    region = "EU"
  }
}
