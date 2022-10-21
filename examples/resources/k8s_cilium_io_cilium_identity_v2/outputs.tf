output "resources" {
  value = {
    "minimal" = k8s_cilium_io_cilium_identity_v2.minimal.yaml
  }
}
