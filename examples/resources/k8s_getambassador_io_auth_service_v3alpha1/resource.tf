resource "k8s_getambassador_io_auth_service_v3alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
