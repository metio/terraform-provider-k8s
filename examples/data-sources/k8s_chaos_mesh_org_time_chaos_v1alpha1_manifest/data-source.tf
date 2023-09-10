data "k8s_chaos_mesh_org_time_chaos_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    mode        = "random-max-percent"
    selector    = {}
    time_offset = "5s"
  }
}
