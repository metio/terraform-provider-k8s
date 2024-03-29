output "manifests" {
  value = {
    "example" = data.k8s_everest_percona_com_backup_storage_v1alpha1_manifest.example.yaml
  }
}
