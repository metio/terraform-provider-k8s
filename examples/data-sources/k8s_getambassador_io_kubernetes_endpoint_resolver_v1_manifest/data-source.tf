data "k8s_getambassador_io_kubernetes_endpoint_resolver_v1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
