data "k8s_iam_services_k8s_aws_policy_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
