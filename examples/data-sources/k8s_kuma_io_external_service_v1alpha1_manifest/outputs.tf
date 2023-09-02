output "manifests" {
  value = {
    "example" = data.k8s_kuma_io_external_service_v1alpha1_manifest.example.yaml
  }
}
