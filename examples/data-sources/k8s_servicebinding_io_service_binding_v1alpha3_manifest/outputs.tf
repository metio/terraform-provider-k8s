output "manifests" {
  value = {
    "example" = data.k8s_servicebinding_io_service_binding_v1alpha3_manifest.example.yaml
  }
}
