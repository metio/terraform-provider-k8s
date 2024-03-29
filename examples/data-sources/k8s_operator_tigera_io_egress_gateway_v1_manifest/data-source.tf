data "k8s_operator_tigera_io_egress_gateway_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
