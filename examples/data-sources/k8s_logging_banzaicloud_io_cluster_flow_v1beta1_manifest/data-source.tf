data "k8s_logging_banzaicloud_io_cluster_flow_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
