data "k8s_kyverno_io_cluster_admission_report_v1alpha2_manifest" "example" {
  metadata = {
    name = "some-name"
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
