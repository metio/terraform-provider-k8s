output "manifests" {
  value = {
    "example" = data.k8s_scheduling_volcano_sh_pod_group_v1beta1_manifest.example.yaml
  }
}
