data "k8s_wgpolicyk8s_io_policy_report_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
