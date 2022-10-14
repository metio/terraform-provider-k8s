output "resources" {
  value = {
    "minimal" = k8s_agent_k8s_elastic_co_agent_v1alpha1.minimal.yaml
    "example" = k8s_agent_k8s_elastic_co_agent_v1alpha1.example.yaml
  }
}
