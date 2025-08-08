output "manifests" {
  value = {
    "example" = data.k8s_netbird_io_nb_routing_peer_v1_manifest.example.yaml
  }
}
