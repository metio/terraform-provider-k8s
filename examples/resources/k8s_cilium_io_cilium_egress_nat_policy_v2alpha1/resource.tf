resource "k8s_cilium_io_cilium_egress_nat_policy_v2alpha1" "minimal" {
  metadata = {
    name = "test"
  }
}
