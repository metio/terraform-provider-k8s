resource "k8s_security_profiles_operator_x_k8s_io_profile_binding_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    image = "some-image"
    profile_ref = {
      kind = "SeccompProfile"
      name = "some-name"
    }
  }
}
