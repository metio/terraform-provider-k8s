output "resources" {
  value = {
    "minimal" = k8s_events_k8s_io_event_v1.minimal.yaml
    "example" = k8s_events_k8s_io_event_v1.example.yaml
  }
}
