data "k8s_nfd_kubernetes_io_node_feature_discovery_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
