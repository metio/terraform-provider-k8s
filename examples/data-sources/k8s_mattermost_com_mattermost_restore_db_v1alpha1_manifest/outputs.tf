output "manifests" {
  value = {
    "example" = data.k8s_mattermost_com_mattermost_restore_db_v1alpha1_manifest.example.yaml
  }
}
