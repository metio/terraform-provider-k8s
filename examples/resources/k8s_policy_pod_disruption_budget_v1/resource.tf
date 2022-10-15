resource "k8s_policy_pod_disruption_budget_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_policy_pod_disruption_budget_v1" "example" {
  metadata = {
    name = "test"
  }
  spec = {
    max_unavailable = "20%"

    selector = {
      match_labels = {
        test = "MyExampleApp"
      }
    }
  }
}
