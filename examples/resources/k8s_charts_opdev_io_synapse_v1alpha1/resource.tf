resource "k8s_charts_opdev_io_synapse_v1alpha1" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
}
