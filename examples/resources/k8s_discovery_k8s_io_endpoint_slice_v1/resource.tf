resource "k8s_discovery_k8s_io_endpoint_slice_v1" "minimal" {
  metadata = {
    name = "test"
  }
  address_type = "IPv4"
  endpoints    = []
}

resource "k8s_discovery_k8s_io_endpoint_slice_v1" "example" {
  metadata = {
    name = "example-abc"
    labels = {
      "kubernetes.io/service-name" = "example"
    }
  }
  address_type = "IPv4"
  ports = [
    {
      name     = "http"
      protocol = "TCP"
      port     = 80
    }
  ]
  endpoints = [
    {
      hostname  = "pod-1"
      nodeName  = "node-1"
      zone      = "us-west2-a"
      addresses = ["10.1.2.3"]
      conditions = {
        ready = true
      }
    }
  ]
}
