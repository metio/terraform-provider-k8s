data "k8s_autoscaling_k8s_io_vertical_pod_autoscaler_checkpoint_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
