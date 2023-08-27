data "k8s_events_k8s_io_event_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
