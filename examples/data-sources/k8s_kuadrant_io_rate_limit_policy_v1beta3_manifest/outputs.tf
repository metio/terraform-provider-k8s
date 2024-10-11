output "manifests" {
  value = {
    "example" = data.k8s_kuadrant_io_rate_limit_policy_v1beta3_manifest.example.yaml
  }
}
