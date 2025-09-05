output "manifests" {
  value = {
    "example" = data.k8s_operator_victoriametrics_com_vl_cluster_v1_manifest.example.yaml
  }
}
