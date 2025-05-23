output "manifests" {
  value = {
    "example" = data.k8s_clickhouse_keeper_altinity_com_click_house_keeper_installation_v1_manifest.example.yaml
  }
}
