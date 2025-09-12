data "k8s_stackconfigpolicy_k8s_elastic_co_stack_config_policy_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
