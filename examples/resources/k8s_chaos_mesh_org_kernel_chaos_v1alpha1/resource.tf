resource "k8s_chaos_mesh_org_kernel_chaos_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    fail_kern_request = {
      failtype = 2
    }
    mode     = "fixed"
    selector = {}
  }
}
