resource "k8s_getambassador_io_host_v2" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
