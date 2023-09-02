output "manifests" {
  value = {
    "example" = data.k8s_crd_projectcalico_org_network_policy_v1_manifest.example.yaml
  }
}
