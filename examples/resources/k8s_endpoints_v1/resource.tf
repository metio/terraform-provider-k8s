resource "k8s_endpoints_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_endpoints_v1" "example" {
  metadata = {
    name = "test"
  }
  subsets = [
    {
      addresses = [
        {
          ip = "10.0.0.4"
        },
        {
          ip = "10.0.0.5"
        },
      ]

      ports = [
        {
          name     = "http"
          port     = 80
          protocol = "TCP"
        },
        {
          name     = "https"
          port     = 443
          protocol = "TCP"
        },
      ]
    },
  ]
}
