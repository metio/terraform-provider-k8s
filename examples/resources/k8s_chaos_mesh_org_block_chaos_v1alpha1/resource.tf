resource "k8s_chaos_mesh_org_block_chaos_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    action      = "delay"
    mode        = "one"
    selector    = {}
    volume_name = "some-name"
  }
}
