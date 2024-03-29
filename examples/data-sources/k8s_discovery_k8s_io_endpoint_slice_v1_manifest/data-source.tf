data "k8s_discovery_k8s_io_endpoint_slice_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  address_type = "IPv4"
  endpoints    = []
}
