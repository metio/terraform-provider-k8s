resource "k8s_chaos_mesh_org_physical_machine_chaos_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    action = "network-corrupt"
    mode   = "fixed-percent"
  }
}
