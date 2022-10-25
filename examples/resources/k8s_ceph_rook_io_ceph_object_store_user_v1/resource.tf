resource "k8s_ceph_rook_io_ceph_object_store_user_v1" "minimal" {
  metadata = {
    name = "test"
  }
  spec = {}
}
