resource "k8s_discovery_k8s_io_endpoint_slice_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
