output "resources" {
  value = {
    "minimal" = k8s_crd_projectcalico_org_ip_reservation_v1.minimal.yaml
  }
}
