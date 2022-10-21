resource "k8s_kyverno_io_admission_report_v1alpha2" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    owner = {
      api_version = "some-version"
      kind        = "some-kind"
      name        = "some-name"
      uid         = "some-uid"
    }
  }
}
