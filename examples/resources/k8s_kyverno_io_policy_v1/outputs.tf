output "resources" {
  value = {
    "minimal" = k8s_kyverno_io_policy_v1.minimal.yaml
  }
}
