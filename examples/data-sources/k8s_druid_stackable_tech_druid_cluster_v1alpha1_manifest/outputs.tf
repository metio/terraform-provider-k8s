output "manifests" {
  value = {
    "example" = data.k8s_druid_stackable_tech_druid_cluster_v1alpha1_manifest.example.yaml
  }
}
