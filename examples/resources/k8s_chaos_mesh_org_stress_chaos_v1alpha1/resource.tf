resource "k8s_chaos_mesh_org_stress_chaos_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    mode     = "all"
    selector = {}
  }
}
