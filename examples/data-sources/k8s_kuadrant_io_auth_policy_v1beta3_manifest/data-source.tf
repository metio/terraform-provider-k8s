data "k8s_kuadrant_io_auth_policy_v1beta3_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
