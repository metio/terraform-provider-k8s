output "resources" {
  value = {
    "minimal" = k8s_litmuschaos_io_chaos_result_v1alpha1.minimal.yaml
  }
}
