data "k8s_getambassador_io_dev_portal_v2" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
