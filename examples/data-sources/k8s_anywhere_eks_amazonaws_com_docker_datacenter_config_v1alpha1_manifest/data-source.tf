data "k8s_anywhere_eks_amazonaws_com_docker_datacenter_config_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
