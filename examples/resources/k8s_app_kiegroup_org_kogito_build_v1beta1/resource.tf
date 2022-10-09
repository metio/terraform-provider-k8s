resource "k8s_app_kiegroup_org_kogito_build_v1beta1" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
  spec = {
    type = "RemoteSource"
  }
}
