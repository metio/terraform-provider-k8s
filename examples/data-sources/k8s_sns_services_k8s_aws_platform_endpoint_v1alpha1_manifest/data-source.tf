data "k8s_sns_services_k8s_aws_platform_endpoint_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
