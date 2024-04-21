output "manifests" {
  value = {
    "example" = data.k8s_gateway_networking_k8s_io_grpc_route_v1_manifest.example.yaml
  }
}
