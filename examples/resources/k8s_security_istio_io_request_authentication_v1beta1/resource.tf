resource "k8s_security_istio_io_request_authentication_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}
