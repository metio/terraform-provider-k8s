resource "k8s_chaos_mesh_org_workflow_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    entry     = "some-entry"
    templates = []
  }
}
