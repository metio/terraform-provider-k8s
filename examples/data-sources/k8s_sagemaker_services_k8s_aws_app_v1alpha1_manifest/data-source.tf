data "k8s_sagemaker_services_k8s_aws_app_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
