output "manifests" {
  value = {
    "example" = data.k8s_crd_projectcalico_org_ipam_block_v1_manifest.example.yaml
  }
}
