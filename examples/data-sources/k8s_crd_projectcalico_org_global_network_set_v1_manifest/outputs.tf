output "manifests" {
  value = {
    "example" = data.k8s_crd_projectcalico_org_global_network_set_v1_manifest.example.yaml
  }
}
