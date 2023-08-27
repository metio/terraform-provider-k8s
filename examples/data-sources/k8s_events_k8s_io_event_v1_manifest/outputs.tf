output "manifests" {
  value = {
    "example" = data.k8s_events_k8s_io_event_v1_manifest.example.yaml
  }
}
