output "manifests" {
  value = {
    "example" = data.k8s_mariadb_persistentsys_monitor_v1alpha1_manifest.example.yaml
  }
}
