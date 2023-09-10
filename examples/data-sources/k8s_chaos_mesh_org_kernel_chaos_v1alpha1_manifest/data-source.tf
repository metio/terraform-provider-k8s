data "k8s_chaos_mesh_org_kernel_chaos_v1alpha1_manifest" "example" {
  metadata = {
    name      = "some-name"
    namespace = "some-namespace"
  }
  spec = {
    fail_kern_request = {
      failtype = 2
    }
    mode     = "fixed"
    selector = {}
  }
}
