data "k8s_logging_banzaicloud_io_flow_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}