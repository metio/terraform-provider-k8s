output "resources" {
  value = {
    "minimal" = k8s_crd_projectcalico_org_cluster_information_v1.minimal.yaml
  }
}
