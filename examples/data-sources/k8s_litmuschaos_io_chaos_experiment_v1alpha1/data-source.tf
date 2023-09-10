data "k8s_litmuschaos_io_chaos_experiment_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
