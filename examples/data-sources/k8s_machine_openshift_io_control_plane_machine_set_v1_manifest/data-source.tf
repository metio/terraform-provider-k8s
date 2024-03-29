data "k8s_machine_openshift_io_control_plane_machine_set_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
