data "k8s_kyverno_io_global_context_entry_v2alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {}
}
