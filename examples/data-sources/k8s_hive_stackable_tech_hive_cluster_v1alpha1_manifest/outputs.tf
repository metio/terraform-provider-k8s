output "manifests" {
  value = {
    "example" = data.k8s_hive_stackable_tech_hive_cluster_v1alpha1_manifest.example.yaml
  }
}
