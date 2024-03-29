data "k8s_security_profiles_operator_x_k8s_io_selinux_profile_v1alpha2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
