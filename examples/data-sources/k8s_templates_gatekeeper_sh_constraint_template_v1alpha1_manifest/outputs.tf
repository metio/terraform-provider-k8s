output "manifests" {
  value = {
    "example" = data.k8s_templates_gatekeeper_sh_constraint_template_v1alpha1_manifest.example.yaml
  }
}
