output "resources" {
  value = {
    "minimal" = k8s_policy_pod_disruption_budget_v1.minimal.yaml
    "example" = k8s_policy_pod_disruption_budget_v1.example.yaml
  }
}
