resource "k8s_chaos_mesh_org_time_chaos_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    mode        = "random-max-percent"
    selector    = {}
    time_offset = "5s"
  }
}
