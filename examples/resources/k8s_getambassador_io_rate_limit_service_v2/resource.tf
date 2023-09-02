resource "k8s_getambassador_io_rate_limit_service_v2" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
