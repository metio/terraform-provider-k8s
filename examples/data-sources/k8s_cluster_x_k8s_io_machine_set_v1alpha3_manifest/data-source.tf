data "k8s_cluster_x_k8s_io_machine_set_v1alpha3_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
