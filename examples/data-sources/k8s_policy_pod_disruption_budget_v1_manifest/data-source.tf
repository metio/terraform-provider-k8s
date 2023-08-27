data "k8s_policy_pod_disruption_budget_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
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
