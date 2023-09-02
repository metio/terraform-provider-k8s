output "manifests" {
  value = {
    "example" = data.k8s_scylla_scylladb_com_scylla_operator_config_v1alpha1_manifest.example.yaml
  }
}
