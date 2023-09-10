data "k8s_getambassador_io_auth_service_v2" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
