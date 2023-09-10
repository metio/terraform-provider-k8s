data "k8s_gateway_solo_io_matchable_http_gateway_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
