data "k8s_kuadrant_io_rate_limit_policy_v1beta2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
