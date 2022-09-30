resource "k8s_traefik_containo_us_ingress_route_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    routes = []
  }
}
