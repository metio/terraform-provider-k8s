resource "k8s_security_istio_io_peer_authentication_v1beta1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
