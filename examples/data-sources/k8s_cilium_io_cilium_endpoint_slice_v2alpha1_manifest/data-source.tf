data "k8s_cilium_io_cilium_endpoint_slice_v2alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
  }
  endpoints = []
}
