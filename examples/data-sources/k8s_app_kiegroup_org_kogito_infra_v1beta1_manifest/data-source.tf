data "k8s_app_kiegroup_org_kogito_infra_v1beta1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
