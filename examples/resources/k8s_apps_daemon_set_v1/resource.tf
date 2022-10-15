resource "k8s_apps_daemon_set_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_apps_daemon_set_v1" "example" {
  metadata = {
    name      = "test"
    namespace = "something"
    labels = {
      test = "MyExampleApp"
    }
  }
  spec = {
    selector = {
      match_labels = {
        test = "MyExampleApp"
      }
    }
    template = {
      metadata = {
        labels = {
          test = "MyExampleApp"
        }
      }
      spec = {
        containers = [
          {
            image = "nginx:1.21.6"
            name  = "example"

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

          }
        ]
      }
    }
  }
}
