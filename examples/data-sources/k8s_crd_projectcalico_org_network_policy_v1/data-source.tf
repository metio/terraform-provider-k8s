data "k8s_crd_projectcalico_org_network_policy_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
