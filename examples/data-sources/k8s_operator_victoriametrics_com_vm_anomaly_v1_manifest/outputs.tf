output "manifests" {
  value = {
    "example" = data.k8s_operator_victoriametrics_com_vm_anomaly_v1_manifest.example.yaml
  }
}
