data "k8s_getambassador_io_consul_resolver_v1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
