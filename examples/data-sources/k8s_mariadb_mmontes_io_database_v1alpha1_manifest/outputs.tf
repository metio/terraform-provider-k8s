output "manifests" {
  value = {
    "example" = data.k8s_mariadb_mmontes_io_database_v1alpha1_manifest.example.yaml
  }
}
