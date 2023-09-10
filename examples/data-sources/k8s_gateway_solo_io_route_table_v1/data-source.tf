data "k8s_gateway_solo_io_route_table_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
