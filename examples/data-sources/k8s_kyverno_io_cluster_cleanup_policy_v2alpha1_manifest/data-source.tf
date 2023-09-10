data "k8s_kyverno_io_cluster_cleanup_policy_v2alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
}
