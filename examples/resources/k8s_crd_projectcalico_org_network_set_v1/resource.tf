resource "k8s_crd_projectcalico_org_network_set_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
