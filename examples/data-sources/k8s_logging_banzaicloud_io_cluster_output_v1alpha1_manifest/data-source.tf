data "k8s_logging_banzaicloud_io_cluster_output_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
