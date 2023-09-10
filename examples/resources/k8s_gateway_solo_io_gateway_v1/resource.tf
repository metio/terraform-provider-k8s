resource "k8s_gateway_solo_io_gateway_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
