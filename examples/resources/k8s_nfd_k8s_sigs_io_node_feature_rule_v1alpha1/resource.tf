resource "k8s_nfd_k8s_sigs_io_node_feature_rule_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    rules = []
  }
}
