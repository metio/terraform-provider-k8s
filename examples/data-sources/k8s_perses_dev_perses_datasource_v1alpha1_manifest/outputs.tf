output "manifests" {
  value = {
    "example" = data.k8s_perses_dev_perses_datasource_v1alpha1_manifest.example.yaml
  }
}
