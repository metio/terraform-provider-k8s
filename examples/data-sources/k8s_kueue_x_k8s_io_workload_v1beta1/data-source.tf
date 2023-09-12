data "k8s_kueue_x_k8s_io_workload_v1beta1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
