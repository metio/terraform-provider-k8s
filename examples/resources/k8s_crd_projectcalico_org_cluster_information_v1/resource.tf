resource "k8s_crd_projectcalico_org_cluster_information_v1" "minimal" {
  metadata = {
    name      = "test"
    namespace = "some-namespace"
  }
}
