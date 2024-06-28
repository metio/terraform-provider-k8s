output "manifests" {
  value = {
    "example" = data.k8s_karpenter_sh_node_claim_v1_manifest.example.yaml
  }
}
