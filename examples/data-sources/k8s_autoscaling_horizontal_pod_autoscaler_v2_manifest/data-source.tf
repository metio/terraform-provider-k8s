data "k8s_autoscaling_horizontal_pod_autoscaler_v2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    min_replicas = 50
    max_replicas = 100

    scale_target_ref = {
      kind = "Deployment"
      name = "MyApp"
    }

    metrics = [
      {
        type = "External"
        external = {
          metric = {
            name = "latency"
            selector = {
              match_labels = {
                lb_name = "test"
              }
            }
          }
          target = {
            type  = "Value"
            value = "100"
          }
        }
      },
    ]
  }
}
