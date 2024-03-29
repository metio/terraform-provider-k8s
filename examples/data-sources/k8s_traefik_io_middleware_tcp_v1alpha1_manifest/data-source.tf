data "k8s_traefik_io_middleware_tcp_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {

  }
}
