data "k8s_anywhere_eks_amazonaws_com_node_upgrade_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
