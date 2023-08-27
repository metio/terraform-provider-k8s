data "k8s_networking_k8s_io_network_policy_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
