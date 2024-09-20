output "manifests" {
  value = {
    "example" = data.k8s_hbase_stackable_tech_hbase_cluster_v1alpha1_manifest.example.yaml
  }
}
