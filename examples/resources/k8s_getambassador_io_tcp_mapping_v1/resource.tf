resource "k8s_getambassador_io_tcp_mapping_v1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
