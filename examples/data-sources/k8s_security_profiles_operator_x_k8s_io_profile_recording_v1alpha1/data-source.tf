data "k8s_security_profiles_operator_x_k8s_io_profile_recording_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
