data "k8s_app_kiegroup_org_kogito_build_v1beta1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
