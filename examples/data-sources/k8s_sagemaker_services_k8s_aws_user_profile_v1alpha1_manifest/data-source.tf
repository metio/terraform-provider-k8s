data "k8s_sagemaker_services_k8s_aws_user_profile_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
