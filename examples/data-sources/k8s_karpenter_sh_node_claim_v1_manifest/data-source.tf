data "k8s_karpenter_sh_node_claim_v1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {
    node_class_ref = {
      group = "some-group"
      kind  = "some-kind"
      name  = "some-name"
    }
    requirements = []
  }
}
