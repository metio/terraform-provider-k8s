resource "k8s_crd_projectcalico_org_felix_configuration_v1" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
}
