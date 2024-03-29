data "k8s_anywhere_eks_amazonaws_com_cloud_stack_datacenter_config_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
