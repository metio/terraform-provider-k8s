data "k8s_nfd_k8s_sigs_io_node_feature_group_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
