output "resources" {
  value = {
    "minimal" = k8s_mutations_gatekeeper_sh_modify_set_v1.minimal.yaml
  }
}
