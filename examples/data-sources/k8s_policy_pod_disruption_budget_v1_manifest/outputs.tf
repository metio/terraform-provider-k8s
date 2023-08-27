output "manifests" {
  value = {
    "example" = data.k8s_policy_pod_disruption_budget_v1_manifest.example.yaml
  }
}
