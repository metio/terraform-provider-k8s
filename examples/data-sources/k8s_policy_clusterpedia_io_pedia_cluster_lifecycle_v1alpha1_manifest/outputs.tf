output "manifests" {
  value = {
    "example" = data.k8s_policy_clusterpedia_io_pedia_cluster_lifecycle_v1alpha1_manifest.example.yaml
  }
}
