data "k8s_cloudfront_services_k8s_aws_cache_policy_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
