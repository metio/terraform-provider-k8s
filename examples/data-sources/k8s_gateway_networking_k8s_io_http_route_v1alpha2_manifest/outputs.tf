output "manifests" {
  value = {
    "example" = data.k8s_gateway_networking_k8s_io_http_route_v1alpha2_manifest.example.yaml
  }
}
