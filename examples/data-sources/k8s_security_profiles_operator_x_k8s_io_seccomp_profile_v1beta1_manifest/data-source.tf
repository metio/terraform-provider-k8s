data "k8s_security_profiles_operator_x_k8s_io_seccomp_profile_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
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
