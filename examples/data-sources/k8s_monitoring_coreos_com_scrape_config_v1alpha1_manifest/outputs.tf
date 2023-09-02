output "manifests" {
  value = {
    "example" = data.k8s_monitoring_coreos_com_scrape_config_v1alpha1_manifest.example.yaml
  }
}
