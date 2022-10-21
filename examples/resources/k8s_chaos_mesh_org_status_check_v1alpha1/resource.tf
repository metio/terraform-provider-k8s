resource "k8s_chaos_mesh_org_status_check_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    type = "HTTP"
  }
}
