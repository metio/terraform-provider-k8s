data "k8s_netbird_io_nb_routing_peer_v1_manifest" "example" {
  metadata = {
    name = "some-name"
    namespace = "some-namespace"
  }
}
