output "manifests" {
  value = {
    "example" = data.k8s_perses_dev_perses_dashboard_v1alpha2_manifest.example.yaml
  }
}
