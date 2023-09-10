resource "k8s_autoscaling_k8s_io_vertical_pod_autoscaler_v1beta2" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
