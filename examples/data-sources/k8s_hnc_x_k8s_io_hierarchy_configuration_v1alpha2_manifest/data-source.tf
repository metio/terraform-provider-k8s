data "k8s_hnc_x_k8s_io_hierarchy_configuration_v1alpha2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
