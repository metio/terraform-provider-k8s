data "k8s_crd_projectcalico_org_staged_kubernetes_network_policy_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
