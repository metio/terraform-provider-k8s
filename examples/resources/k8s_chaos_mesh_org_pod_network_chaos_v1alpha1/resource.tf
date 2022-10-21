resource "k8s_chaos_mesh_org_pod_network_chaos_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {}
}
