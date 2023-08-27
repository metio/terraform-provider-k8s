resource "k8s_autoscaling_horizontal_pod_autoscaler_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
