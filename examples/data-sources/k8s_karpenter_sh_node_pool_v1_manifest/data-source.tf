data "k8s_karpenter_sh_node_pool_v1_manifest" "example" {
  metadata = {
    name = "some-name"

  }
}
