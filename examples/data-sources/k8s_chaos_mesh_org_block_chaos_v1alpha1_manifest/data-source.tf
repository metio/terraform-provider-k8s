data "k8s_chaos_mesh_org_block_chaos_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    action      = "delay"
    mode        = "one"
    selector    = {}
    volume_name = "some-name"
  }
}
