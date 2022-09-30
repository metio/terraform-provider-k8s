resource "k8s_security_istio_io_peer_authentication_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {

  }
}
