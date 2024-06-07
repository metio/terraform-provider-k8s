data "k8s_chainsaw_kyverno_io_test_v1alpha2_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {
    steps = []
  }
}
