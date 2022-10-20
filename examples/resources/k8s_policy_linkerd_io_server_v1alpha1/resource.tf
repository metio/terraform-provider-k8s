resource "k8s_policy_linkerd_io_server_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    pod_selector = {}
    port         = 12345
  }
}
