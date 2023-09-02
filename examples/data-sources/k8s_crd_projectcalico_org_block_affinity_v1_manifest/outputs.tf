output "manifests" {
  value = {
    "example" = data.k8s_crd_projectcalico_org_block_affinity_v1_manifest.example.yaml
  }
}
