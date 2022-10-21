resource "k8s_chaos_mesh_org_schedule_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    schedule = "some-schedule"
    type     = "some-type"
  }
}
