data "k8s_kyverno_io_cluster_background_scan_report_v1alpha2_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {}
}
