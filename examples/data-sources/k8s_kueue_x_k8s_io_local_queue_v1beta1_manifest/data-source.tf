data "k8s_kueue_x_k8s_io_local_queue_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
