output "manifests" {
  value = {
    "example" = data.k8s_clickhouse_altinity_com_click_house_operator_configuration_v1_manifest.example.yaml
  }
}
