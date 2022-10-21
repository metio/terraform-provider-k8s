resource "k8s_chaos_mesh_org_workflow_node_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    start_time    = "2022-10-21"
    template_name = "some-name"
    type          = "some-type"
    workflow_name = "some-workflow"
  }
}
