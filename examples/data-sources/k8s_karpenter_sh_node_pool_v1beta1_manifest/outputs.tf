output "manifests" {
  value = {
    "example" = data.k8s_karpenter_sh_node_pool_v1beta1_manifest.example.yaml
  }
}
