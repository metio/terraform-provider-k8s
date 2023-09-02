output "manifests" {
  value = {
    "example" = data.k8s_config_grafana_com_project_config_v1_manifest.example.yaml
  }
}
