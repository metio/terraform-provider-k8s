resource "k8s_networking_k8s_io_ingress_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_networking_k8s_io_ingress_v1" "example" {
  metadata = {
    name = "test"
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
