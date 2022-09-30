resource "k8s_crd_projectcalico_org_calico_node_status_v1" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
}
