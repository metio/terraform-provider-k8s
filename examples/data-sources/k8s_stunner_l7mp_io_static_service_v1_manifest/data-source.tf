data "k8s_stunner_l7mp_io_static_service_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
