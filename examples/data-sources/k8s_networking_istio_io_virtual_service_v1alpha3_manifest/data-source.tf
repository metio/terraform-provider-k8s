data "k8s_networking_istio_io_virtual_service_v1alpha3_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
