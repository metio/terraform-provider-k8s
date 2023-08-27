data "k8s_policy_pod_disruption_budget_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
