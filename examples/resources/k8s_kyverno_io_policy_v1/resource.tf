resource "k8s_kyverno_io_policy_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {}
}
