output "manifests" {
  value = {
    "example" = data.k8s_everest_percona_com_monitoring_config_v1alpha1_manifest.example.yaml
  }
}
