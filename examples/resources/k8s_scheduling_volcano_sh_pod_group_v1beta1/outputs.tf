output "resources" {
  value = {
    "minimal" = k8s_scheduling_volcano_sh_pod_group_v1beta1.minimal.yaml
  }
}
