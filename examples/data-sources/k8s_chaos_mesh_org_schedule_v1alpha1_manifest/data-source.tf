data "k8s_chaos_mesh_org_schedule_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    schedule = "some-schedule"
    type     = "some-type"
  }
}
