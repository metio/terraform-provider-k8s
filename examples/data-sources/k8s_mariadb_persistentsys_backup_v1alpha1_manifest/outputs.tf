output "manifests" {
  value = {
    "example" = data.k8s_mariadb_persistentsys_backup_v1alpha1_manifest.example.yaml
  }
}
