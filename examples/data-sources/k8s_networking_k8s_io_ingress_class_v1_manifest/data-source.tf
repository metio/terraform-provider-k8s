data "k8s_networking_k8s_io_ingress_class_v1_manifest" "example" {
  metadata = {
    name = "some-name"
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
