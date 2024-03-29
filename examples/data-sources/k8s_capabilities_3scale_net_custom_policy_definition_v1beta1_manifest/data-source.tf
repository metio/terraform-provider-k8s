data "k8s_capabilities_3scale_net_custom_policy_definition_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
