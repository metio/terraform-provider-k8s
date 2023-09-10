data "k8s_eks_services_k8s_aws_fargate_profile_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
