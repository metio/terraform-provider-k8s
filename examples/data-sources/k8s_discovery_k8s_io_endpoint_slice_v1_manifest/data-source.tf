data "k8s_discovery_k8s_io_endpoint_slice_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
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
