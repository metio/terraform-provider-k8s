output "manifests" {
  value = {
    "example" = data.k8s_everest_percona_com_database_cluster_restore_v1alpha1_manifest.example.yaml
  }
}
