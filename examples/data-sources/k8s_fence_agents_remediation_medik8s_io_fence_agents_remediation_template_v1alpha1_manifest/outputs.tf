output "manifests" {
  value = {
    "example" = data.k8s_fence_agents_remediation_medik8s_io_fence_agents_remediation_template_v1alpha1_manifest.example.yaml
  }
}
