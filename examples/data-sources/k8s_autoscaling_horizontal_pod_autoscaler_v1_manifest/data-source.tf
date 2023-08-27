data "k8s_autoscaling_horizontal_pod_autoscaler_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    max_replicas = 10
    min_replicas = 8

    scale_target_ref = {
      kind = "Deployment"
      name = "MyApp"
    }
  }
}
