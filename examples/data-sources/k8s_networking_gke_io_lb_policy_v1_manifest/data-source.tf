data "k8s_networking_gke_io_lb_policy_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    target_ref = {
      group = "some-group"
      kind  = "some-kind"
      name  = "some-name"
    }
  }
}
