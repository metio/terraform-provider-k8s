output "resources" {
  value = {
    "minimal" = k8s_chaos_mesh_org_stress_chaos_v1alpha1.minimal.yaml
  }
}
