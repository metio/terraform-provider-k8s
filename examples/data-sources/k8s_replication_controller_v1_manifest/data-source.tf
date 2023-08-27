data "k8s_replication_controller_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
    labels = {
      test = "MyExampleApp"
    }
  }
  spec = {
    selector = {
      test = "MyExampleApp"
    }
    template = {
      metadata = {
        labels = {
          test = "MyExampleApp"
        }
        annotations = {
          "key1" = "value1"
        }
      }
      spec = {
        containers = [
          {
            image = "nginx:1.21.6"
            name  = "example"

            liveness_probe = {
              http_get = {
                path = "/"
                port = 80

                http_header = {
                  name  = "X-Custom-Header"
                  value = "Awesome"
                }
              }

              initial_delay_seconds = 3
              period_seconds        = 3
            }

            resources = {
              limits = {
                cpu    = "0.5"
                memory = "512Mi"
              }
              requests = {
                cpu    = "250m"
                memory = "50Mi"
              }
            }
          }
        ]
      }
    }
  }
}
