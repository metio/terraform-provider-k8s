output "manifests" {
  value = {
    "example" = data.k8s_hive_openshift_io_selector_sync_set_v1_manifest.example.yaml
  }
}
