output "manifests" {
  value = {
    "example" = data.k8s_kyverno_io_cleanup_policy_v2_manifest.example.yaml
  }
}
