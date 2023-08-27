resource "k8s_networking_k8s_io_ingress_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
