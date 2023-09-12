data "k8s_operator_cluster_x_k8s_io_core_provider_v1alpha2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
