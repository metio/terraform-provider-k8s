data "k8s_route53resolver_services_k8s_aws_resolver_rule_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
