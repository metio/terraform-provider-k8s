output "resources" {
  value = {
    "minimal" = k8s_discovery_k8s_io_endpoint_slice_v1.minimal.yaml
    "example" = k8s_discovery_k8s_io_endpoint_slice_v1.example.yaml
  }
}
