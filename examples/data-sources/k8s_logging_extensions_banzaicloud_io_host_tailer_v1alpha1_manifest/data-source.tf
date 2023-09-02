data "k8s_logging_extensions_banzaicloud_io_host_tailer_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
