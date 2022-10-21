resource "k8s_chaos_mesh_org_gcp_chaos_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    action   = "disk-loss"
    instance = "some-instance"
    project  = "some-project"
    zone     = "some-zone"
  }
}
