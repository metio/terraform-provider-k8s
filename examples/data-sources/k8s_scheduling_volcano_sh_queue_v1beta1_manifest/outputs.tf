output "manifests" {
  value = {
    "example" = data.k8s_scheduling_volcano_sh_queue_v1beta1_manifest.example.yaml
  }
}
