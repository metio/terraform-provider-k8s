resource "k8s_kyverno_io_cleanup_policy_v2alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
