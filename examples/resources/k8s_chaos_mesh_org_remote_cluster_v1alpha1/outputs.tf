output "resources" {
  value = {
    "minimal" = k8s_chaos_mesh_org_remote_cluster_v1alpha1.minimal.yaml
  }
}
