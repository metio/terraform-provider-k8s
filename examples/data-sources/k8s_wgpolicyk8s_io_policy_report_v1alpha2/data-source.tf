data "k8s_wgpolicyk8s_io_policy_report_v1alpha2" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
