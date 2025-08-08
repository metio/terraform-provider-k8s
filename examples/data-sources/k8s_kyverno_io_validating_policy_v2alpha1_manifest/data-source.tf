data "k8s_kyverno_io_validating_policy_v2alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    
  }
}
