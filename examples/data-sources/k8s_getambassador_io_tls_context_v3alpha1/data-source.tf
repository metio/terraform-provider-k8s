data "k8s_getambassador_io_tls_context_v3alpha1" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
