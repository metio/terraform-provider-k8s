data "k8s_getambassador_io_rate_limit_service_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
