data "k8s_actions_summerwind_dev_horizontal_runner_autoscaler_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
