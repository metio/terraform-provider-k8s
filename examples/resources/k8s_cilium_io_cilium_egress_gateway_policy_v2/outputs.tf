output "resources" {
  value = {
    "minimal" = k8s_cilium_io_cilium_egress_gateway_policy_v2.minimal.yaml
  }
}
