data "k8s_operator_tigera_io_tls_pass_through_route_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
