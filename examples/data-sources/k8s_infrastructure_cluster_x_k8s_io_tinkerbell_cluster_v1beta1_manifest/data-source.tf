data "k8s_infrastructure_cluster_x_k8s_io_tinkerbell_cluster_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
