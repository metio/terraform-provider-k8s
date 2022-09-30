output "resources" {
  value = {
    "minimal" = k8s_templates_gatekeeper_sh_constraint_template_v1.minimal.yaml
  }
}
