output "manifests" {
  value = {
    "example" = data.k8s_operator_victoriametrics_com_v_logs_v1beta1_manifest.example.yaml
  }
}
