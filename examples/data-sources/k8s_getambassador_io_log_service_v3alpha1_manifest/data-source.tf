data "k8s_getambassador_io_log_service_v3alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
