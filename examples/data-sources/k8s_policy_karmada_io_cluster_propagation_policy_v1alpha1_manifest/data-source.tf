data "k8s_policy_karmada_io_cluster_propagation_policy_v1alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  spec = {
    resource_selectors = []
  }
}
