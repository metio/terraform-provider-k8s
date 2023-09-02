data "k8s_getambassador_io_consul_resolver_v3alpha1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
    
  }
}
