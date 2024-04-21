data "k8s_operator_tigera_io_tls_terminated_route_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
