output "manifests" {
  value = {
    "example" = data.k8s_hive_openshift_io_hive_config_v1_manifest.example.yaml
  }
}
