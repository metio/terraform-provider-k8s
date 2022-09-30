output "resources" {
  value = {
    "minimal" = k8s_crd_projectcalico_org_bgp_peer_v1.minimal.yaml
  }
}
