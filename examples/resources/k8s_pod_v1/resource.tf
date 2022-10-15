resource "k8s_pod_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_pod_v1" "example" {
  metadata = {
    name = "terraform-example"
  }
  spec = {
    containers = [
      {
        image = "nginx:1.21.6"
        name  = "example"
        env = [
          {
            name  = "environment"
            value = "test"
          }
        ]
        ports = [
          {
            container_port = 80
          }
        ]
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
      },
    ]
    dns_config = {
      nameservers = ["1.1.1.1", "8.8.8.8", "9.9.9.9"]
      searches    = ["example.com"]
      options = [
        {
          name  = "ndots"
          value = 1
        },
        {
          name = "use-vc"
        },
      ]
      dns_policy = "None"
    }
  }
}
