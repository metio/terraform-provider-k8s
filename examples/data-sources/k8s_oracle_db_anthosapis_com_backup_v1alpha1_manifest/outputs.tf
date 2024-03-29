output "manifests" {
  value = {
    "example" = data.k8s_oracle_db_anthosapis_com_backup_v1alpha1_manifest.example.yaml
  }
}
