output "resources" {
  value = {
    "minimal" = k8s_mutations_gatekeeper_sh_assign_v1.minimal.yaml
  }
}
