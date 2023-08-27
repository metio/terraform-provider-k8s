data "k8s_networking_k8s_io_ingress_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    default_backend = {
      service = {
        name = "myapp-1"
        port = {
          number = 8080
        }
      }
    }
    rules = [
      {
        http = {
          paths = [
            {
              backend = {
                service = {
                  name = "myapp-1"
                  port = {
                    number = 8080
                  }
                }
              }
              path      = "/app1/*"
              path_type = "Prefix"
            },
            {
              backend = {
                service = {
                  name = "myapp-2"
                  port = {
                    number = 8080
                  }
                }
              }
              path      = "/app2/*"
              path_type = "Prefix"
            },
          ]
        }
      }
    ]
    tls = [
      {
        secret_name = "tls-secret"
      },
    ]
  }
}
