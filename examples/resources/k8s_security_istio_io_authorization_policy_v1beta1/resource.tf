resource "k8s_security_istio_io_authorization_policy_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
}
