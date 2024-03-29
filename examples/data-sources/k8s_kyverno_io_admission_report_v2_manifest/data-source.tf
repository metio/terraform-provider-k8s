data "k8s_kyverno_io_admission_report_v2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
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
