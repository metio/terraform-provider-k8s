data "k8s_policy_karmada_io_override_policy_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {}
}
