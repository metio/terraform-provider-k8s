resource "k8s_security_profiles_operator_x_k8s_io_security_profile_node_status_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  node_name = "some-name"
}
