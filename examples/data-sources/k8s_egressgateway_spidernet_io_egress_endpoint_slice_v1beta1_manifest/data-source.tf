data "k8s_egressgateway_spidernet_io_egress_endpoint_slice_v1beta1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
