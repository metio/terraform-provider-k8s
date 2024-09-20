output "manifests" {
  value = {
    "example" = data.k8s_superset_stackable_tech_druid_connection_v1alpha1_manifest.example.yaml
  }
}
