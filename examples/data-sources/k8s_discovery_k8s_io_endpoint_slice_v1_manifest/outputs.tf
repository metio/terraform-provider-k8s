output "manifests" {
  value = {
    "example" = data.k8s_discovery_k8s_io_endpoint_slice_v1_manifest.example.yaml
  }
}
