data "k8s_operator_cluster_x_k8s_io_infrastructure_provider_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
