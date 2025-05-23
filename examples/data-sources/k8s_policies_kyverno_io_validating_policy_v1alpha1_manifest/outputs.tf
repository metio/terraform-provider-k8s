output "manifests" {
  value = {
    "example" = data.k8s_policies_kyverno_io_validating_policy_v1alpha1_manifest.example.yaml
  }
}
