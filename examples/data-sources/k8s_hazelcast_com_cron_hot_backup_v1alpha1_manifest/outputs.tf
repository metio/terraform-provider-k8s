output "manifests" {
  value = {
    "example" = data.k8s_hazelcast_com_cron_hot_backup_v1alpha1_manifest.example.yaml
  }
}
