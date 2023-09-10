data "k8s_getambassador_io_consul_resolver_v2" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
