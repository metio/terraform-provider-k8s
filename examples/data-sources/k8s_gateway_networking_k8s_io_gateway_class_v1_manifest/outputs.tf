output "manifests" {
  value = {
    "example" = data.k8s_gateway_networking_k8s_io_gateway_class_v1_manifest.example.yaml
  }
}
