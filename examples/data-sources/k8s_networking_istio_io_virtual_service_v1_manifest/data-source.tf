data "k8s_networking_istio_io_virtual_service_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
