output "manifests" {
  value = {
    "example" = data.k8s_groupsnapshot_storage_k8s_io_volume_group_snapshot_class_v1beta2_manifest.example.yaml
  }
}
