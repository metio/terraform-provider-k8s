resource "k8s_getambassador_io_mapping_v2" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
