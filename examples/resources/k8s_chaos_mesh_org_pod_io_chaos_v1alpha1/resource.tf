resource "k8s_chaos_mesh_org_pod_io_chaos_v1alpha1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {
    volume_mount_path = "/"
  }
}
