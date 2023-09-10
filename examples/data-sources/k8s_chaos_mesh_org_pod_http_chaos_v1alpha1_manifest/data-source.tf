data "k8s_chaos_mesh_org_pod_http_chaos_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
}
