resource "k8s_sagemaker_services_k8s_aws_domain_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
