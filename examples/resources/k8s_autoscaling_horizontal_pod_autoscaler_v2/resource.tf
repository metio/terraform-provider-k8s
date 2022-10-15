resource "k8s_autoscaling_horizontal_pod_autoscaler_v2" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_autoscaling_horizontal_pod_autoscaler_v2" "example" {
  metadata = {
    name = "test"
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
