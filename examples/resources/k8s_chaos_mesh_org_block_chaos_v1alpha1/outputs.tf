output "resources" {
  value = {
    "minimal" = k8s_chaos_mesh_org_block_chaos_v1alpha1.minimal.yaml
  }
}
