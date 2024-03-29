output "manifests" {
  value = {
    "example" = data.k8s_snapscheduler_backube_snapshot_schedule_v1_manifest.example.yaml
  }
}
