data "k8s_application_networking_k8s_aws_iam_auth_policy_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
