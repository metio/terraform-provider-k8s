data "k8s_cilium_io_cilium_network_policy_v2_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
