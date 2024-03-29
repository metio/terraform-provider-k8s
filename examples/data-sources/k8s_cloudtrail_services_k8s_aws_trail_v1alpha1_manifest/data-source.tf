data "k8s_cloudtrail_services_k8s_aws_trail_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
