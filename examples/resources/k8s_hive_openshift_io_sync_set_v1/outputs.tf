output "resources" {
  value = {
    "minimal" = k8s_hive_openshift_io_sync_set_v1.minimal.yaml
  }
}
