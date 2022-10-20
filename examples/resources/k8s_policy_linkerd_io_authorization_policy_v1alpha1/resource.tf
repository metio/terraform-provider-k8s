resource "k8s_policy_linkerd_io_authorization_policy_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    required_authentication_refs = []
    target_ref = {
      kind = "some-kind"
      name = "some-name"
    }
  }
}
