resource "k8s_traefik_io_servers_transport_tcp_v1alpha1" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"

  }
}
