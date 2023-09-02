output "manifests" {
  value = {
    "example" = data.k8s_networking_istio_io_gateway_v1beta1_manifest.example.yaml
  }
}
