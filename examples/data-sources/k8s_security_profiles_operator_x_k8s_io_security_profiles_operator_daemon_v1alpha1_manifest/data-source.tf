data "k8s_security_profiles_operator_x_k8s_io_security_profiles_operator_daemon_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
