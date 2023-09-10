data "k8s_sagemaker_services_k8s_aws_transform_job_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
