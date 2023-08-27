data "k8s_events_k8s_io_event_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  type   = "Warning"
  reason = "Some not so urgent event has occurred"
  related = {
    kind = "some-kind"
  }
  event_time = "2022-10-16"
}
