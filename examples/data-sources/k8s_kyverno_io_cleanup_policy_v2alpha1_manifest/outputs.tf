output "manifests" {
  value = {
    "example" = data.k8s_kyverno_io_cleanup_policy_v2alpha1_manifest.example.yaml
  }
}
