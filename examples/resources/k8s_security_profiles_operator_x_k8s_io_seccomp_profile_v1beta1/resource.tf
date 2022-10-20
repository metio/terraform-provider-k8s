resource "k8s_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    default_action = "SCMP_ACT_TRAP"
  }
}

resource "k8s_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1" "example" {
  metadata = {
    name      = "profile1"
    namespace = "my-namespace"
  }
  spec = {
    default_action    = "SCMP_ACT_ERRNO"
    base_profile_name = "runc-v1.0.0"
    syscalls = [
      {
        action = "SCMP_ACT_ALLOW"
        names  = ["exit_group"]
      }
    ]
  }
}
