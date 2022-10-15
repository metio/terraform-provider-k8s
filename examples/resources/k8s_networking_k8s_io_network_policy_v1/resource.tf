resource "k8s_networking_k8s_io_network_policy_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_networking_k8s_io_network_policy_v1" "example" {
  metadata = {
    name = "test"
  }
  spec = {
    pod_selector = {
      match_expressions = [
        {
          key      = "name"
          operator = "In"
          values   = ["webfront", "api"]
        },
      ]
    }
    ingress = [
      {
        ports = [
          {
            port     = "http"
            protocol = "TCP"
          },
          {
            port     = "8125"
            protocol = "UDP"
          },
        ]

        from = [
          {
            namespace_selector = {
              match_labels = {
                name = "default"
              }
            }
          },
          {
            ip_block = {
              cidr = "10.0.0.0/8"
              except = [
                "10.0.0.0/24",
                "10.0.1.0/24",
              ]
            }
          },
        ]
      }
    ]
    egress       = []
    policy_types = ["Ingress", "Egress"]
  }
}
