resource "k8s_chaos_mesh_org_pod_http_chaos_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
}
