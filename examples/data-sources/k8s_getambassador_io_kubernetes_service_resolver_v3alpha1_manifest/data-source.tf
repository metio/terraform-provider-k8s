data "k8s_getambassador_io_kubernetes_service_resolver_v3alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
