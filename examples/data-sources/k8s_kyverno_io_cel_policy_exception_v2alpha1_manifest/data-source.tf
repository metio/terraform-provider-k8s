data "k8s_kyverno_io_cel_policy_exception_v2alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
