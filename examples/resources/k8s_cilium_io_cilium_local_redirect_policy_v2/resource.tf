resource "k8s_cilium_io_cilium_local_redirect_policy_v2" "minimal" {
  metadata = {
    name = "test"
  }
}
