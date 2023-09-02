output "manifests" {
  value = {
    "example" = data.k8s_servicebinding_io_service_binding_v1beta1_manifest.example.yaml
  }
}
