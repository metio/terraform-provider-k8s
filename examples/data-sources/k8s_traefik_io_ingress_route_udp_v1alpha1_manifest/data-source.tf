data "k8s_traefik_io_ingress_route_udp_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    routes = []
  }
}
