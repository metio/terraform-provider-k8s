output "manifests" {
  value = {
    "example" = data.k8s_security_istio_io_peer_authentication_v1beta1_manifest.example.yaml
  }
}
