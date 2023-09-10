data "k8s_litmuschaos_io_chaos_result_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
