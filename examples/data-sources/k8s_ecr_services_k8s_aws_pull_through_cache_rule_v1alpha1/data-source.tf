data "k8s_ecr_services_k8s_aws_pull_through_cache_rule_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
