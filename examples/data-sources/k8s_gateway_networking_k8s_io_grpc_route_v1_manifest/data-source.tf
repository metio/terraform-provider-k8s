data "k8s_gateway_networking_k8s_io_grpc_route_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}