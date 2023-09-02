data "k8s_telemetry_istio_io_telemetry_v1alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
