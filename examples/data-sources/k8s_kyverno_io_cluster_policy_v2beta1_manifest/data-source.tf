data "k8s_kyverno_io_cluster_policy_v2beta1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {}
}
