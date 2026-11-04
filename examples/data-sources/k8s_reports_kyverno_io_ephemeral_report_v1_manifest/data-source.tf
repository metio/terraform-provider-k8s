data "k8s_reports_kyverno_io_ephemeral_report_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }

  spec = {
    owner = {
      api_version = "some-value"
      kind        = "some-value"
      name        = "some-value"
      uid         = "some-value"
    }
  }
}
