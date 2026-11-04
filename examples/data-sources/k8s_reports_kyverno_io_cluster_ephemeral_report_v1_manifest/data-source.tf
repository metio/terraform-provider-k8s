data "k8s_reports_kyverno_io_cluster_ephemeral_report_v1_manifest" "example" {
  metadata = {
    name = "some-name"

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
