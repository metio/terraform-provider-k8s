output "manifests" {
  value = {
    "example" = data.k8s_policy_clusterpedia_io_cluster_import_policy_v1alpha1_manifest.example.yaml
  }
}
