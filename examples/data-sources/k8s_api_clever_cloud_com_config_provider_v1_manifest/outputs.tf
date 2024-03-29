output "manifests" {
  value = {
    "example" = data.k8s_api_clever_cloud_com_config_provider_v1_manifest.example.yaml
  }
}
