resource "k8s_events_k8s_io_event_v1" "minimal" {
  metadata = {
    name = "test"
  }
  event_time = "2022-10-16"
}

resource "k8s_events_k8s_io_event_v1" "example" {
  metadata = {
    name = "test"
  }
  type   = "Warning"
  reason = "Some not so urgent event has occurred"
  related = {
    kind = "some-kind"
  }
  event_time = "2022-10-16"
}
