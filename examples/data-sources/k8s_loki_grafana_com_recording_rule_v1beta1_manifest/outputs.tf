output "manifests" {
  value = {
    "example" = data.k8s_loki_grafana_com_recording_rule_v1beta1_manifest.example.yaml
  }
}
