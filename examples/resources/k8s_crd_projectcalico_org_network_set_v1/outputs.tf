output "resources" {
  value = {
    "minimal" = k8s_crd_projectcalico_org_network_set_v1.minimal.yaml
  }
}
