resource "k8s_kyverno_io_background_scan_report_v1alpha2" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {}
}
