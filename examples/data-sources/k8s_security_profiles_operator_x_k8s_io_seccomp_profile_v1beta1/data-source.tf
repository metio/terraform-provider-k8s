data "k8s_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
