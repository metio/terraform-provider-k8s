output "manifests" {
  value = {
    "example" = data.k8s_service_account_v1_manifest.example.yaml
  }
}
