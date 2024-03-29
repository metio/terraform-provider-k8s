output "manifests" {
  value = {
    "example" = data.k8s_tempo_grafana_com_tempo_stack_v1alpha1_manifest.example.yaml
  }
}
