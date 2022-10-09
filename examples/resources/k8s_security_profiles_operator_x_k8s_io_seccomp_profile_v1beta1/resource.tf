resource "k8s_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    default_action = "SCMP_ACT_TRAP"
  }
}
