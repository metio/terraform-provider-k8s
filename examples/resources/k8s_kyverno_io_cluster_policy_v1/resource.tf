resource "k8s_kyverno_io_cluster_policy_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {}
}
