output "manifests" {
  value = {
    "example" = data.k8s_policies_kyverno_io_policy_exception_v1alpha1_manifest.example.yaml
  }
}
