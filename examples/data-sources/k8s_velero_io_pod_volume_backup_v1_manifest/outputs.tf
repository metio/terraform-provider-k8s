output "manifests" {
  value = {
    "example" = data.k8s_velero_io_pod_volume_backup_v1_manifest.example.yaml
  }
}
