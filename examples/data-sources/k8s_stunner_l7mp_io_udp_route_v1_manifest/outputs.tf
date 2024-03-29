output "manifests" {
  value = {
    "example" = data.k8s_stunner_l7mp_io_udp_route_v1_manifest.example.yaml
  }
}
