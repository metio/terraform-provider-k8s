data "k8s_gateway_nginx_org_client_settings_policy_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    target_ref = {
      kind  = "Service"
      group = "v1"
      name  = "some-service"
    }
  }
}
