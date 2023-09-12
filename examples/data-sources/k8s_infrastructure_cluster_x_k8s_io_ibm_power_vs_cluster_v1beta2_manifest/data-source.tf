data "k8s_infrastructure_cluster_x_k8s_io_ibm_power_vs_cluster_v1beta2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
