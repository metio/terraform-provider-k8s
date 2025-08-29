data "k8s_experimental_kubeblocks_io_node_count_scaler_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
