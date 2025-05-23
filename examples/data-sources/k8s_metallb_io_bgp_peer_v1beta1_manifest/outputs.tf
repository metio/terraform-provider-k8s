output "manifests" {
  value = {
    "example" = data.k8s_metallb_io_bgp_peer_v1beta1_manifest.example.yaml
  }
}
