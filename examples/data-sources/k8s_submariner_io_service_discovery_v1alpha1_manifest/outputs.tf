output "manifests" {
  value = {
    "example" = data.k8s_submariner_io_service_discovery_v1alpha1_manifest.example.yaml
  }
}
