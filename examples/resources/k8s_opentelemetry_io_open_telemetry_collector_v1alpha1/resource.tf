resource "k8s_opentelemetry_io_open_telemetry_collector_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
