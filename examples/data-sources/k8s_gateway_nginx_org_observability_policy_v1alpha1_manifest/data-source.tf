data "k8s_gateway_nginx_org_observability_policy_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    target_ref = {
      group = "v1"
      kind  = "Service"
      name  = "some-service"
    }
  }
}
