output "resources" {
  value = {
    "minimal" = k8s_sematext_com_sematext_agent_v1.minimal.yaml
    "example" = k8s_sematext_com_sematext_agent_v1.example.yaml
  }
}
