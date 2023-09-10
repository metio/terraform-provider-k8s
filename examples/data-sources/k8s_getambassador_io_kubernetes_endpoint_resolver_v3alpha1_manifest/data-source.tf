data "k8s_getambassador_io_kubernetes_endpoint_resolver_v3alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
