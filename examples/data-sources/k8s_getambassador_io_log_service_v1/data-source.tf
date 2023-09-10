data "k8s_getambassador_io_log_service_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
