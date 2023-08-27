output "manifests" {
  value = {
    "example" = data.k8s_secret_v1_manifest.example.yaml
  }
}
