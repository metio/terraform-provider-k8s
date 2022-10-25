resource "k8s_opentelemetry_io_open_telemetry_collector_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
}
