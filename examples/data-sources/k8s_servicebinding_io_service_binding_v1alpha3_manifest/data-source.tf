data "k8s_servicebinding_io_service_binding_v1alpha3_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
