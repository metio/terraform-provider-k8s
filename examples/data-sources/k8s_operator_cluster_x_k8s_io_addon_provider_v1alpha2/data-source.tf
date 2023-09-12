data "k8s_operator_cluster_x_k8s_io_addon_provider_v1alpha2" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
