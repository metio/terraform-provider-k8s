resource "k8s_networking_k8s_io_ingress_class_v1" "minimal" {
  metadata = {
    name = "test"
  }
}

resource "k8s_networking_k8s_io_ingress_class_v1" "example" {
  metadata = {
    name = "test"
  }
  spec = {
    controller = "example.com/ingress-controller"
    parameters = {
      api_group = "k8s.example.com"
      kind      = "IngressParameters"
      name      = "external-lb"
    }
  }
}
