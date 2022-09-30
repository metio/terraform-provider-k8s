resource "k8s_crd_projectcalico_org_global_network_set_v1" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
}
