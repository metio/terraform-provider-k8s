data "k8s_kyverno_io_policy_v2beta1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
