data "k8s_metacontroller_k8s_io_controller_revision_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
