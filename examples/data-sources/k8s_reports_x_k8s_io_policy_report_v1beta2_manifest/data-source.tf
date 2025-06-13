data "k8s_reports_x_k8s_io_policy_report_v1beta2_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
