output "resources" {
  value = {
    "minimal" = k8s_crd_projectcalico_org_ipam_config_v1.minimal.yaml
  }
}
