output "manifests" {
  value = {
    "example" = data.k8s_velero_io_volume_snapshot_location_v1_manifest.example.yaml
  }
}
