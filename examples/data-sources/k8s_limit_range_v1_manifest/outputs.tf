output "manifests" {
  value = {
    "example" = data.k8s_limit_range_v1_manifest.example.yaml
  }
}
