data "k8s_capabilities_3scale_net_developer_user_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
