output "manifests" {
  value = {
    "example" = data.k8s_nifi_stackable_tech_nifi_cluster_v1alpha1_manifest.example.yaml
  }
}
