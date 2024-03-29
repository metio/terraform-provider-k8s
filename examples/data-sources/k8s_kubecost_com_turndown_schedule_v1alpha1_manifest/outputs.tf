output "manifests" {
  value = {
    "example" = data.k8s_kubecost_com_turndown_schedule_v1alpha1_manifest.example.yaml
  }
}
