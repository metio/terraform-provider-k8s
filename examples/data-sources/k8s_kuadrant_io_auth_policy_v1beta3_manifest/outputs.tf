output "manifests" {
  value = {
    "example" = data.k8s_kuadrant_io_auth_policy_v1beta3_manifest.example.yaml
  }
}
