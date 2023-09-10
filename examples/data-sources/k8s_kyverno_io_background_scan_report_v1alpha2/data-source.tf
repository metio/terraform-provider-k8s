data "k8s_kyverno_io_background_scan_report_v1alpha2" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
