data "k8s_prometheusservice_services_k8s_aws_rule_groups_namespace_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
