resource "k8s_chaos_mesh_org_dns_chaos_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    action   = "error"
    mode     = "fixed"
    selector = {}
  }
}
