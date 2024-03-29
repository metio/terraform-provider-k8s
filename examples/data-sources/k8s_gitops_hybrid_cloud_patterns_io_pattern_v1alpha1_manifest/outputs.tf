output "manifests" {
  value = {
    "example" = data.k8s_gitops_hybrid_cloud_patterns_io_pattern_v1alpha1_manifest.example.yaml
  }
}
