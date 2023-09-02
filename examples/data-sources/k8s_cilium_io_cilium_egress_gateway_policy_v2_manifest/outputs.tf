output "manifests" {
  value = {
    "example" = data.k8s_cilium_io_cilium_egress_gateway_policy_v2_manifest.example.yaml
  }
}
