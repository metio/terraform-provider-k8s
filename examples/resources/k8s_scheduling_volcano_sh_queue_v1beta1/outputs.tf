output "resources" {
  value = {
    "minimal" = k8s_scheduling_volcano_sh_queue_v1beta1.minimal.yaml
  }
}
