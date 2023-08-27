output "manifests" {
  value = {
    "example" = data.k8s_namespace_v1_manifest.example.yaml
  }
}
