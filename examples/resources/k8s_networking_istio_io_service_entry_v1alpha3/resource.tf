resource "k8s_networking_istio_io_service_entry_v1alpha3" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
