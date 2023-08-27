output "manifests" {
  value = {
    "example" = data.k8s_endpoints_v1_manifest.example.yaml
  }
}
