output "manifests" {
  value = {
    "example" = data.k8s_hdfs_stackable_tech_hdfs_cluster_v1alpha1_manifest.example.yaml
  }
}
