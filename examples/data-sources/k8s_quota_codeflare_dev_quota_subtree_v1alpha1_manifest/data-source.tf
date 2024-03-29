data "k8s_quota_codeflare_dev_quota_subtree_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
