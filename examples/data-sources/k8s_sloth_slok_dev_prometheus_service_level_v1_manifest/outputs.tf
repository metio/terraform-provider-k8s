output "manifests" {
  value = {
    "example" = data.k8s_sloth_slok_dev_prometheus_service_level_v1_manifest.example.yaml
  }
}
