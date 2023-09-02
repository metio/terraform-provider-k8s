output "manifests" {
  value = {
    "example" = data.k8s_loki_grafana_com_loki_stack_v1_manifest.example.yaml
  }
}
