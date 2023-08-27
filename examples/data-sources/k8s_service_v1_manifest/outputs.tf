output "manifests" {
  value = {
    "example" = data.k8s_service_v1_manifest.example.yaml
  }
}
