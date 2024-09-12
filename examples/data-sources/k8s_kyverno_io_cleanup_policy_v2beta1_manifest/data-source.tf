data "k8s_kyverno_io_cleanup_policy_v2beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    schedule = "some-schedule"
  }
}
