resource "k8s_policy_karmada_io_propagation_policy_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
