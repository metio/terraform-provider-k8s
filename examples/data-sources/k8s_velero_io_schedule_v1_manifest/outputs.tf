output "manifests" {
  value = {
    "example" = data.k8s_velero_io_schedule_v1_manifest.example.yaml
  }
}
