resource "k8s_autoscaling_k8s_io_vertical_pod_autoscaler_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    target_ref = {
      kind = "Deployment"
      name = "some-name"
    }
  }
}
