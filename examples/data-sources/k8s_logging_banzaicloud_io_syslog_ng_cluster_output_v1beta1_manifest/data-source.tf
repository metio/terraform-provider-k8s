data "k8s_logging_banzaicloud_io_syslog_ng_cluster_output_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {}
}
