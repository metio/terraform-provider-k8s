output "manifests" {
  value = {
    "example" = data.k8s_trino_stackable_tech_trino_cluster_v1alpha1_manifest.example.yaml
  }
}
