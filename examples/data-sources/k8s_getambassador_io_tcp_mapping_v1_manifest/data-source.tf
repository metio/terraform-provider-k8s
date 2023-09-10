data "k8s_getambassador_io_tcp_mapping_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
