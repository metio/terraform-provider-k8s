output "manifests" {
  value = {
    "example" = data.k8s_operator_victoriametrics_com_vl_single_v1_manifest.example.yaml
  }
}
