output "manifests" {
  value = {
    "example" = data.k8s_apiregistration_k8s_io_api_service_v1_manifest.example.yaml
  }
}
