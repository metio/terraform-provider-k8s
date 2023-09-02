resource "k8s_networking_istio_io_sidecar_v1alpha3" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
