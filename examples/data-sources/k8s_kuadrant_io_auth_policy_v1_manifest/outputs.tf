output "manifests" {
  value = {
    "example" = data.k8s_kuadrant_io_auth_policy_v1_manifest.example.yaml
  }
}
