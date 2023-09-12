data "k8s_nfd_kubernetes_io_node_feature_rule_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
