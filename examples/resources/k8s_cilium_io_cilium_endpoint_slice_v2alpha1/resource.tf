resource "k8s_cilium_io_cilium_endpoint_slice_v2alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  endpoints = []
}
